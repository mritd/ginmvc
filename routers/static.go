package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
)

// register static file router
// static file like *.js、*.css...
// we use packr tool to embed static files into go binaries
// packr tool refs => https://github.com/gobuffalo/packr/tree/master/v2
func init() {
	register("static", func(router *gin.Engine) {
		// embed static files into go binaries
		staticBox := packr.New("static", "../static")
		router.StaticFS("/static", staticBox)
	})
}
