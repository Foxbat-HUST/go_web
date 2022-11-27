package test

import (
	"bytes"
	"fmt"
	"go_web/api/handler"
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

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

var actorMap = map[string]entity.User{
	"admin": {
		ID:   1,
		Name: "admin",
		Type: entity.UserTypeSuper,
	},
}

type handlerSignature func(db *gorm.DB, cfg *config.Config) func(ctx *gin.Context)

var handlerFactorMap = map[string]map[Method]handlerSignature{
	"/api/v1/users/:id": {
		GET:    handler.InitGetUserHandler,
		PUT:    handler.InitUpdateUserHandler,
		DELETE: handler.InitDeleteUserHandler,
	},
}

func GetApp() App {
	cfg := config.LoadConfigForTest()
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

func getAppWithTx() App {
	app := GetApp()
	return App{
		Config: app.Config,
		Db:     app.Db.Begin(),
	}
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

type Request struct {
	Url  string
	Body []byte
}
type Scenario struct {
	Name         string
	Actor        string
	Request      Request
	BeforeTest   func(app App)
	ExpectedCode int
	AssertFunc   func(t *testing.T, app App, response *httptest.ResponseRecorder)
	AfterTest    func(app App)
}
type TestData struct {
	Path      string
	Method    Method
	Scenarios []Scenario
}

func RunTest(t *testing.T, test TestData) {

	if len(test.Scenarios) == 0 {
		return
	}

	appWithTx := getAppWithTx()
	defer appWithTx.Db.Rollback()

	// create route with middleware & handler
	router := gin.Default()
	router.Use(createLoginMiddleware(appWithTx))
	groupApi := router.Group(test.Path)
	handlerFactorGrp, ok := handlerFactorMap[test.Path]
	if !ok {
		panic(fmt.Sprintf("could not found handler for path: %s", test.Path))
	}
	handlerFactor, ok := handlerFactorGrp[test.Method]
	if !ok {
		panic(fmt.Sprintf("could not found handler for path %s:%s", test.Method, test.Path))
	}
	handler := handlerFactor(appWithTx.Db, appWithTx.Config)
	switch test.Method {
	case GET:
		groupApi.GET("", handler)
	case POST:
		groupApi.POST("", handler)
	case PUT:
		groupApi.PUT("", handler)
	case DELETE:
		groupApi.DELETE("", handler)
	default:
		panic(fmt.Sprintf("un-support http method: %s", test.Method))
	}

	// do test
	for _, item := range test.Scenarios {
		if item.BeforeTest != nil {
			item.BeforeTest(appWithTx)
		}
		request, err := http.NewRequest(string(test.Method), item.Request.Url, bytes.NewBuffer(item.Request.Body))
		request.Header.Set("Content-Type", "application/json")
		if err != nil {
			panic(err)
		}
		doLogin(appWithTx, item.Actor, request)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, request)

		assert.Equal(t, item.ExpectedCode, w.Code, item.Name)

		if item.AssertFunc != nil {
			item.AssertFunc(t, appWithTx, w)
		}

		if item.AfterTest != nil {
			item.AfterTest(appWithTx)
		}
	}

}
