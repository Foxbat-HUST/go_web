package handler

import (
	"go_web/api/usecase/user"
	"go_web/domain/service/implement"
	"go_web/infra/repository/mysql"
	"go_web/test"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	appWithTx := test.GetAppWithTx()
	defer appWithTx.Db.Rollback()

	userRepo := mysql.NewUserRepo(appWithTx.Db)
	userService := implement.NewUserService(userRepo)
	getUserUc := user.NewGetUser(appWithTx.Db, userService)
	test.RunTest(t, appWithTx, test.TestData{
		Path:    "/api/v1/users/:id",
		Method:  test.GET,
		Handler: GetUser(getUserUc),
		Scenarios: []test.Scenario{
			{
				Name:  "OK test get user by id",
				Actor: "admin",
				Request: test.Request{
					Url: "/api/v1/users/1",
				},
				ExpectedCode: 200,
				AssertFunc: func(t *testing.T, resRecorder *httptest.ResponseRecorder) {

				},
			},
		},
	})

}
