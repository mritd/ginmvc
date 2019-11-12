package routers

import (
	"github.com/gin-gonic/gin"
)

// @Tags Monitor API
// @Summary health check
// @Description health check
// @Produce json
// @Success 200 {object} models.CommonResp "{"message":"success","timestamp":1570889247}"
// @Router /healthz [get]
func init() {
	register("healthCheck", func(router *gin.Engine) {
		router.GET("/healthz", func(c *gin.Context) {
			c.JSON(200, success())
		})
	})
}
