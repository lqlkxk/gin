package initRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/lqlkxk/gin/handler"
	"github.com/lqlkxk/gin/middleware"
	"net/http"
)

//路由初始化
func SteupRouter() *gin.Engine {

	router := gin.Default()
	// 添加自定义的 logger 中间件
	//router.Use(middleware.Logger(), gin.Recovery())

	// 设置模板地址
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "welcome")
	})

	//设置用户路由组
	userGroup := router.Group("/user")
	{
		// 设备为path传参
		userGroup.GET("findByName/:username", middleware.JWTAuth(), handler.UserHandler{}.GetByName)
		// 设备为普通传参
		userGroup.GET("findById", middleware.CookieAuth(), handler.UserHandler{}.FindById)
		// post 参数 表单提交
		userGroup.POST("login", handler.UserHandler{}.Login)
		// post 参数 表单提交
		userGroup.POST("save", handler.UserHandler{}.Save)

		// post json参数 表单提交
		userGroup.POST("saveByJson", handler.UserHandler{}.SaveJson)

	}

	//设置用户路由组
	webGroup := router.Group("/web")
	{
		// 设备为path传参
		webGroup.GET("/index", handler.IndexHandler{}.Index)
	}
	return router
}
