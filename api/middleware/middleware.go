package middleware

import "github.com/gin-gonic/gin"

type MiddleWare interface {
	Value() func(*gin.Context)
}

type MiddlewareType string

const (
	LoginMiddleware MiddlewareType = "login"
)

func Create(t MiddlewareType) MiddleWare {
	switch t {
	case LoginMiddleware:
		
		return nil
	default:
		return nil
	}
}
