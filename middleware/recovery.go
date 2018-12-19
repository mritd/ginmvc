package middleware

import "github.com/gin-gonic/gin"

func init() {
	registerWithWeight(100, func() gin.HandlerFunc {
		return gin.Recovery()
	})
}
