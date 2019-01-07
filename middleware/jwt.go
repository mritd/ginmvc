package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mritd/ginmvc/auth"
)

// JWT middleware
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(bearerToken, "Bearer ") || len(strings.Fields(bearerToken)) != 2 {
			c.JSON(http.StatusForbidden, gin.H{
				"message":   "invalid jwt token",
				"timestamp": time.Now().Unix(),
			})
			c.Abort()
			return
		}
		token := strings.Fields(bearerToken)[1]

		j := auth.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message":   err.Error(),
				"timestamp": time.Now().Unix(),
			})
			c.Abort()
			return
		}
		c.Set(auth.JWTClaimsKey, claims)
	}
}
