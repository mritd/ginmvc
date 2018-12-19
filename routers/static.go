package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
)

func init() {
	register("static", func(router *gin.Engine) {
		// embed static files into go binaries
		staticBox := packr.NewBox("../static")
		router.StaticFS("/static", staticBox)
	})
}
