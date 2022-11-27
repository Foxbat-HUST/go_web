package middleware

import (
	"go_web/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewLoginMiddleware(authService service.AuthService) MiddleWare {
	return &loginMiddleware{
		authService: authService,
	}
}

type loginMiddleware struct {
	authService service.AuthService
}

func (l loginMiddleware) Value() func(*gin.Context) {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("Authenticate")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}

		user, err := l.authService.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		ctx.Set("authUser", user)
		ctx.Next()
	}
}
