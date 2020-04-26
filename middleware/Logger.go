package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		host := ctx.Request.Host
		url := ctx.Request.URL
		method := ctx.Request.Method
		fmt.Printf("%s\t%s\t%s\t%s ", time.Now().Format("2006-01-02 15:04:05"), host, url, method)
		ctx.Next()
		fmt.Println(ctx.Writer.Status())
	}

}
