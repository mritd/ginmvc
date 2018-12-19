package middleware

import (
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/mritd/ginmvc/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	register(func() gin.HandlerFunc {
		sessionSecret := viper.GetString("basic.session_secret")
		if sessionSecret == "" {
			logrus.Warn("session secret is blank, auto generate...")
			sessionSecret = utils.RandStr(16)
		}

		store := cookie.NewStore([]byte(sessionSecret))
		return sessions.Sessions("basic", store)
	})

}
