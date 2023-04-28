package handler

import (
	"go_web/domain/entity"
	"go_web/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleErr(ctx *gin.Context, e error) {
	error, ok := e.(*errors.Error)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, e.Error())
	} else {
		ctx.JSON(error.StatusCode, error.Error())
	}
}

func setCookie(ctx *gin.Context, name, value string, age int) {
	ctx.SetCookie(name, value, age, "", "", true, true)
}

func getAuthUser(ctx *gin.Context) *entity.User {
	val, exist := ctx.Get("authUser")
	if !exist {
		return nil
	}

	return val.(*entity.User)
}
