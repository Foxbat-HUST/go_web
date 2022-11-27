package handler

import (
	"encoding/json"
	"fmt"
	"go_web/domain/entity"
	"go_web/infra/repository/mysql"
	"go_web/test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	test.RunTest(t, test.TestData{
		Path:   "/api/v1/users/:id",
		Method: test.GET,
		Scenarios: []test.Scenario{
			{
				Name:  "OK test get user by id",
				Actor: "admin",
				Request: test.Request{
					Url: "/api/v1/users/1",
				},
				ExpectedCode: http.StatusOK,
				AssertFunc: func(t *testing.T, app test.App, response *httptest.ResponseRecorder) {
					var resUser entity.User
					if err := json.Unmarshal(response.Body.Bytes(), &resUser); err != nil {
						panic(fmt.Sprintf("could not parse response body %v", response.Body))
					}
					var userId uint32 = 1
					dbUser, err := mysql.NewUserRepo(app.Db).FindByID(userId)
					if err != nil {
						panic(fmt.Sprintf("could not find user with ID = %d from db", userId))
					}
					assert.Equal(t, *dbUser, resUser)
				},
			},
		},
	})
}

func TestCreateUser(t *testing.T) {
	test.RunTest(t, test.TestData{
		Path:   "/api/v1/users",
		Method: test.POST,
		Scenarios: []test.Scenario{
			{
				Name:  "OK test create user",
				Actor: "admin",
				Request: test.Request{
					Url: "/api/v1/users",
					Body: []byte(`
					{
						"name": "anpq",
						"email": "anpq@gmail.com",
						"age": 28
					}
					`),
				},
				ExpectedCode: http.StatusCreated,
				AssertFunc: func(t *testing.T, app test.App, response *httptest.ResponseRecorder) {
					var resUser entity.User
					if err := json.Unmarshal(response.Body.Bytes(), &resUser); err != nil {
						panic(fmt.Sprintf("could not parse response body %v", response.Body))
					}
					dbUser, err := mysql.NewUserRepo(app.Db).FindByID(resUser.ID)
					if err != nil {
						panic(fmt.Sprintf("could not find user with ID = %d from db", resUser.ID))
					}
					assert.Equal(t, dbUser.Name, resUser.Name)
					assert.Equal(t, dbUser.Age, resUser.Age)
					assert.Equal(t, dbUser.Email, resUser.Email)
					assert.Equal(t, entity.UserTypeNormal, resUser.Type)
				},
			},
		},
	})
}
