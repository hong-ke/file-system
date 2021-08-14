package router

import (
	"filesystem/config"
	"filesystem/controller"
	middleware "filesystem/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func addApiRouter(engine *gin.Engine) {
	// 在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，
	// 这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器
	engine.Use(gin.Recovery())
	// 强制浏览器不使用缓存
	engine.Use(middleware.NoCache)
	// 浏览器跨域 OPTIONS 请求设置
	engine.Use(middleware.Options)
	// 一些安全设置
	engine.Use(middleware.Secure)
	// 设置请求ID生成全局中间件
	engine.Use(middleware.RequestID)
	//	统一打印日志
	engine.Use(middleware.Error)

	//Request &Response 相关内容,影响性能看情况开
	if config.GetInstance().GetBool("application.log.request") {
		engine.Use(middleware.Logger)
	}
	// 注册swagger路由
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})

	// 业务API
	{
		v1 := engine.Group("/v1")
		v1.GET("/hello", controller.HelloController)
		v1.POST("/hello", controller.Hello2Controller)
	}

}
