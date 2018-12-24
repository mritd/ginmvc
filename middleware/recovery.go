package middleware

import "github.com/gin-gonic/gin"

// register recovery middleware and the weight must very high
func init() {
	registerWithWeight(100, func() gin.HandlerFunc {
		return gin.Recovery()
	})
}
