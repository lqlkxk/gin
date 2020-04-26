package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexHandler struct {
}

// 模板渲染
func (index IndexHandler) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "hello world",
	})
}
