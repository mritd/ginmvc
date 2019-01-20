/*
 * Copyright 2018 mritd <mritd1234@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
