package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lqlkxk/gin/initRouter"
	"github.com/lqlkxk/gin/utils"
	"github.com/mattn/go-colorable"
)

func main() {
	// 设置为发布模式
	gin.SetMode(gin.ReleaseMode)
	// 解决乱码
	gin.DefaultWriter = colorable.NewColorableStdout()
	router := initRouter.SteupRouter()
	port, _ := utils.GetServer()
	//默认8080 启动
	router.Run(port)
}
