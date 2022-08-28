package handler

import (
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
