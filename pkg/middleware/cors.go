package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

type corsWrapper struct {
	*cors.Cors
	optionPassthrough bool
}

func (c corsWrapper) build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.HandlerFunc(ctx.Writer, ctx.Request)
		if !c.optionPassthrough &&
			ctx.Request.Method == http.MethodOptions &&
			ctx.GetHeader("Access-Control-Request-Method") != "" {
			ctx.AbortWithStatus(http.StatusOK)
		}
	}
}

func AllowAll() gin.HandlerFunc {
	return corsWrapper{Cors: cors.AllowAll()}.build()
}