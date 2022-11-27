package handler

import (
	"go_web/test"
	"net/http/httptest"
	"testing"
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
				ExpectedCode: 200,
				AssertFunc: func(t *testing.T, app test.App, resRecorder *httptest.ResponseRecorder) {
				},
			},
		},
	})

}
