package routers

import (
	_ "github.com/mritd/ginmvc/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin swagger refs => https://github.com/swaggo/gin-swagger
func init() {
	register("swagger", func(router *gin.Engine) {
		url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	})
}
