package router

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var instance *gin.Engine
var once sync.Once

func init() {
	once.Do(func() {
		instance = gin.New()
		addRouter()
	})

}

func addRouter() {
	// 注册路由
	addApiRouter(instance)

}

func GetInstance() *gin.Engine {
	return instance
}
