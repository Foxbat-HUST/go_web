package test

import (
	"bytes"
	"fmt"
	"go_web/api/middleware"
	"go_web/config"
	"go_web/domain/entity"
	"go_web/domain/service/implement"
	repo "go_web/infra/repository/mysql"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

func getApp() App {
	cfg := config.LoadConfig()
	if cfg == nil {
		panic("fail to load config")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Db)
	fmt.Printf("dns: %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return App{
		Config: cfg,
		Db:     db,
	}
}

var actorMap = map[string]entity.User{
	"admin": {
		ID:   1,
		Name: "admin",
		Type: entity.UserTypeSuper,
	},
}

func createLoginMiddleware(app App) func(*gin.Context) {
	userRepo := repo.NewUserRepo(app.Db)
	authService := implement.NewAuthService(app.Config, userRepo)

	return middleware.NewLoginMiddleware(authService).Value()
}

func doLogin(app App, userName string, request *http.Request) {
	user, ok := actorMap[userName]
	if !ok {
		panic(fmt.Sprintf("not found user with name: %s", userName))
	}
	userRepo := repo.NewUserRepo(app.Db)
	authService := implement.NewAuthService(app.Config, userRepo)
	token, err := authService.CreateToken(user)

	if err != nil {
		panic(err)
	}

	request.AddCookie(&http.Cookie{
		Name:  "Authenticate",
		Value: token,
	})
}

type App struct {
	Db     *gorm.DB
	Config *config.Config
}

type Method string

const (
	GET    Method = "GET"
	Post   Method = "POST"
	Put    Method = "PUT"
	Delete Method = "DELETE"
)

type Scenario struct {
	Name         string
	Actor        string
	Request      []byte
	BeforeTest   func(app App)
	ExpectedCode int
	AssertFunc   func(app App, t *testing.T, resRecorder *httptest.ResponseRecorder)
	AfterTest    func(app App)
}
type TestData struct {
	Url       string
	Method    Method
	Handler   func(ctx *gin.Context)
	Scenarios []Scenario
}

func RunTest(t *testing.T, test TestData) {

	if len(test.Scenarios) == 0 {
		return
	}

	var appWithTx = (func() App {
		app := getApp()
		return App{
			Config: app.Config,
			Db:     app.Db.Begin(),
		}
	})()

	defer appWithTx.Db.Rollback()

	// create route with middleware & handler
	router := gin.Default()
	router.Use(createLoginMiddleware(appWithTx))
	group := router.Group(test.Url)
	switch test.Method {
	case GET:
		group.GET(test.Url, test.Handler)
	case Post:
		group.POST(test.Url, test.Handler)
	case Put:
		group.PUT(test.Url, test.Handler)
	case Delete:
		group.DELETE(test.Url, test.Handler)
	default:
		panic(fmt.Sprintf("un-support http method: %s", test.Method))
	}

	for _, item := range test.Scenarios {
		if item.BeforeTest != nil {
			item.BeforeTest(appWithTx)
		}
		request, err := http.NewRequest(string(test.Method), test.Url, bytes.NewBuffer(item.Request))
		if err != nil {
			panic(err)
		}
		doLogin(appWithTx, item.Actor, request)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, request)

		assert.Equal(t, item.ExpectedCode, w.Code, item.Name)

		if item.AssertFunc != nil {
			item.AssertFunc(appWithTx, t, w)
		}

		if item.AfterTest != nil {
			item.AfterTest(appWithTx)
		}
	}

}
