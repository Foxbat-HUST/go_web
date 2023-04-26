package cmd

import (
	"fmt"
	"go_web/api/handler"
	"go_web/api/middleware"
	"go_web/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var runServer = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  `Tbd`,
	Run: func(cmd *cobra.Command, args []string) {
		initServer()
	},
}

type app struct {
	config *config.Config
	db     *gorm.DB
}

func (a app) CreateMiddleware(middlewareType middleware.MiddlewareType) func(*gin.Context) {
	switch middlewareType {
	case middleware.LoginMiddleware:
		authService := handler.InitAuthService(a.db, a.config)
		return middleware.NewLoginMiddleware(authService).Value()
	default:
		panic("un-support type")
	}

}

func initServer() {
	cfg := config.LoadConfig()
	db := _initMysql(cfg)
	app := app{
		config: cfg,
		db:     db,
	}
	_initRouter(app)
}

func _initMysql(cfg *config.Config) *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Db)
	fmt.Printf("dns: %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func _initRouter(app app) {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	auth := router.Group("/auth")
	{
		auth.POST("/login", handler.InitLoginHandler(app.db, app.config))
		auth.GET("/init", handler.InitAuthInitHandler(app.db, app.config))
	}
	apiV1 := router.Group("/api/v1")
	{
		apiV1.Use(app.CreateMiddleware(middleware.LoginMiddleware))
		userGrp := apiV1.Group("/users")
		{
			userGrp.POST("", handler.InitCreateUserHandler(app.db, app.config))
			userGrp.PUT("/:id", handler.InitUpdateUserHandler(app.db, app.config))
			userGrp.GET("/:id", handler.InitGetUserHandler(app.db, app.config))
			userGrp.DELETE("/:id", handler.InitDeleteUserHandler(app.db, app.config))
		}
	}
	router.Run()
}
