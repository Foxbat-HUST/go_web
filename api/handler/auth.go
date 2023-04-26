package handler

import (
	"go_web/api/usecase/auth"
	"go_web/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func doLogin(uc auth.Login, config *config.Config) func(*gin.Context) {
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

		setCookie(ctx, config.Auth.CookiesName, output.Token, config.Auth.TokenExpireSeconds)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "authenticated",
		})
	}
}
func doAuthInit(uc auth.AuthInit, config *config.Config) func(*gin.Context) {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie(config.Auth.CookiesName)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}

		data, err := uc.Exec(auth.AuthInitInput{
			Token: token,
		})
		if err != nil {
			handleErr(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
