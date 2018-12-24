package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
)

// register static file router
// static file like *.js、*.css...
// we use packr tool to embed static files into go binaries
func init() {
	register("static", func(router *gin.Engine) {
		// embed static files into go binaries
		staticBox := packr.NewBox("../static")
		router.StaticFS("/static", staticBox)
	})
}
