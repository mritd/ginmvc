package middleware

import (
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
	"github.com/mritd/ginmvc/auth"
)

func RBACBasicAuth() gin.HandlerFunc {
	return authz.NewAuthorizer(auth.Enforcer)
}
