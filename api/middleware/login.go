package middleware

import (
	"go_web/domain/service"
	"net/http"
	"strings"

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
		reqToken := ctx.Request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		var jwtToken string
		if len(splitToken) >= 2 {
			jwtToken = splitToken[1]
		}

		user, err := l.authService.ParseToken(jwtToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		ctx.Set("authUser", user)
		ctx.Next()
	}
}
