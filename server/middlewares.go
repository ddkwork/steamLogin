package server

import (
	"steamBackend/conf"
	"steamBackend/server/controllers"
	"steamBackend/utils"

	"github.com/gin-gonic/gin"
)

func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		// 同源
		if origin == "" {
			c.Next()
			return
		}
		method := c.Request.Method
		// 设置跨域
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		c.Header("Access-Control-Allow-Headers", "Content-Length,session,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language, Keep-Alive, User-Agent, Cache-Control, Content-Type")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified")
		c.Header("Access-Control-Max-Age", "172800")

		// 信任域名
		if conf.Conf.Server.SiteUrl != "*" && utils.ContainsString(conf.Origins, c.GetHeader("Origin")) == -1 {
			c.JSON(200, controllers.MetaResponse(413, "The origin is not in the site_url list, please configure it correctly."))
			c.Abort()
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(204)
		}
		//处理请求
		c.Next()
	}
}
