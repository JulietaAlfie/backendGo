package middleware

import (
	"errors"
	"os"

	"github.com/JulietaAlfie/backendGo.git/pkg/web"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			ctx.Abort()
			return
		}

		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, 401, errors.New("token not valid"))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
