package server

import (
	"steamBackend/server/controllers"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// 初始化路由
func InitRouter(engine *gin.Engine) {
	log.Infof("初始化路由...")
	engine.Use(CorsHandler())
	InitWSRouter(engine)
}

// 初始化WebSocket路由

func InitWSRouter(engine *gin.Engine) {
	engine.GET("/ws", controllers.Ws)
}
