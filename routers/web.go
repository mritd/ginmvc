package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mritd/ginmvc/auth"
	"github.com/mritd/ginmvc/middleware"
	"github.com/mritd/ginmvc/models"
)

// register web page router
func init() {
	register("web", func(router *gin.Engine) {
		router.GET("", index)
		rbacPublicGroup := router.Group("/rbac")
		rbacPublicGroup.POST("/register", rbacRegister)
		rbacPrivateGroup := router.Group("/rbac")
		rbacPrivateGroup.Use(middleware.RBACSessionAuth)
		rbacPrivateGroup.GET("/test", rbacTest)
	})
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", "Gin MVC")
}

func rbacRegister(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(auth.UserKey)
	u := user.(models.User)
	if !auth.Enforcer.AddPolicy(u.Email, "/rbac/test", "*") {
		c.JSON(500, failed("rbac register failed"))
		return
	} else {
		c.JSON(200, success())
		return
	}
}

func rbacTest(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(auth.UserKey)
	u := user.(models.User)
	c.HTML(200, "rbac.html", u)
}
