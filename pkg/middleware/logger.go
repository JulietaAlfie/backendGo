package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		time := time.Now()
		path := ctx.Request.URL.Path
		verb := ctx.Request.Method

		ctx.Next()

		var size int
		if ctx.Writer != nil {
			size = ctx.Writer.Size()
		}
		fmt.Printf("time: %v \npath: localhost:8080%s\nverb: %s\nsize: %d\n", time, path, verb, size)
	}
}
