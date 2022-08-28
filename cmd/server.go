package cmd

import (
	"go_web/api/handler"
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
	db *gorm.DB
}

func initServer() {
	app := app{
		db: _initMysql(),
	}
	_initRouter(app)
}

func _initMysql() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "user:123456@tcp(127.0.0.1:3306)/web?charset=utf8mb4&parseTime=True&loc=Local"
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
	apiV1 := router.Group("/api/v1")
	{
		userGrp := apiV1.Group("/users")
		{
			userGrp.POST("", handler.CreateUser(initUcCreateUser(app.db)))
			userGrp.PUT("/:id", handler.UpdateUser(initUcUpdateUser(app.db)))
			userGrp.GET("/:id", handler.GetUser(initUcGetUser(app.db)))
			userGrp.DELETE("/:id", handler.DeleteUser(initUcDeleteUser(app.db)))
		}
	}
	router.Run()
}
