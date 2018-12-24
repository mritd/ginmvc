package middleware

import (
	"github.com/gin-contrib/sessions/cookie"
	"github.com/mritd/ginmvc/conf"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/mritd/ginmvc/utils"
	"github.com/sirupsen/logrus"
)

// register session middleware
func init() {
	register(func() gin.HandlerFunc {
		if conf.Basic.SessionSecret == "" {
			logrus.Warn("session secret is blank, auto generate...")
			conf.Basic.SessionSecret = utils.RandStr(16)
		}

		store := cookie.NewStore([]byte(conf.Basic.SessionSecret))
		return sessions.Sessions("ginmvc", store)
	})

}
