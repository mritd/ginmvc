package middleware

import (
	"net/http"
	"time"

	"github.com/mritd/ginmvc/models"

	"github.com/gin-contrib/authz"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mritd/ginmvc/auth"
)

func RBACBasicAuth() gin.HandlerFunc {
	return authz.NewAuthorizer(auth.Enforcer)
}

// RBACSessionAuth is used for user login status check and RBAC permission check
func RBACSessionAuth(c *gin.Context) {
	session := sessions.Default(c)
	u := session.Get(auth.UserKey)
	if u == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message":   "unauthorized access",
			"timestamp": time.Now().Unix(),
		})
		c.Abort()
		return
	}
	user, ok := u.(models.User)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{
			"message":   "unauthorized access",
			"timestamp": time.Now().Unix(),
		})
		c.Abort()
		return
	}

	rbacCheck := auth.Enforcer.Enforce(user.Email, c.Request.URL.Path, c.Request.Method)
	if !rbacCheck {
		c.JSON(http.StatusForbidden, gin.H{
			"message":   "unauthorized access",
			"timestamp": time.Now().Unix(),
		})
		c.Abort()
		return
	}
}
