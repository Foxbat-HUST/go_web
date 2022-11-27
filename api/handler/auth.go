package handler

import (
	"go_web/api/usecase/auth"
	"go_web/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func doLogin(uc auth.Login, config *config.Config) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		input := auth.LoginInput{}
		if err := ctx.BindJSON(&input); err != nil {
			handleErr(ctx, err)
			return
		}
		output, err := uc.Exec(input)
		if err != nil {
			handleErr(ctx, err)
			return
		}

		setCookie(ctx, "Authenticate", output.Token, config.Auth.TokenExpireSeconds)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "authenticated",
		})
	}
}
