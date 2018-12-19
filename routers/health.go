package routers

import (
	"github.com/gin-gonic/gin"
)

func init() {
	register("healthCheck", func(router *gin.Engine) {
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "pong")
		})
	})
}
