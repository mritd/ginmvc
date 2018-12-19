package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const ginLogFormat = "request => %d | %s | %s | %s | %s | %s"

func init() {
	registerWithWeight(100, func() gin.HandlerFunc {
		return Ginrus()
	})
}

func Ginrus() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		if len(c.Errors) > 0 {
			logrus.Error(fmt.Sprintf(ginLogFormat, c.Writer.Status(), c.ClientIP(), c.Request.Method, path, latency, c.Request.UserAgent()), " | ERROR: ", c.Errors.String())
		} else {
			logrus.Info(fmt.Sprintf(ginLogFormat, c.Writer.Status(), c.ClientIP(), c.Request.Method, path, latency, c.Request.UserAgent()))
		}
	}
}
