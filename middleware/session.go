package middleware

import (
	"github.com/gin-contrib/sessions/memstore"
	"net/http"
	"time"

	"github.com/mritd/ginmvc/auth"
	"github.com/mritd/ginmvc/conf"
	"github.com/mritd/ginmvc/models"

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
			conf.Basic.SessionSecret = utils.RandString(16)
		}
		store := memstore.NewStore([]byte(conf.Basic.SessionSecret))
		return sessions.Sessions("ginmvc", store)
	})

}

// SessionAuth used to check user login status
// No RBAC verification required
func SessionAuth(c *gin.Context) {
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
	_, ok := u.(models.User)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{
			"message":   "unauthorized access",
			"timestamp": time.Now().Unix(),
		})
		c.Abort()
		return
	}
}
