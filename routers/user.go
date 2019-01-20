package routers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/mritd/ginmvc/middleware"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"github.com/mritd/ginmvc/auth"
	"github.com/mritd/ginmvc/db"
	"github.com/mritd/ginmvc/models"
	"github.com/mritd/ginmvc/utils"
)

func init() {
	register("user", func(router *gin.Engine) {
		userGroup := router.Group("/user")
		userGroup.POST("/register", userRegister)
		userGroup.POST("/login", userLogin)
		userGroup.POST("/logout", middleware.SessionAuth, userLogout)
	})
}

func userRegister(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, failed(err.Error()))
		return
	}
	user.ID = 0
	user.Salt = utils.RandString(auth.SaltLength)
	user.Lock = false
	user.CreateTime = time.Now().Unix()
	m := md5.New()
	m.Write([]byte(user.Password + user.Salt))
	user.Password = hex.EncodeToString(m.Sum(nil))

	var count int
	err = db.Orm.Model(&models.User{}).Where("email = ?", user.Email).Count(&count).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, failed(err.Error()))
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, failed(fmt.Sprintf("user email [%s] already register", user.Email)))
		return
	}
	err = db.Orm.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, failed(err.Error()))
		return
	}
	c.JSON(http.StatusOK, success())
}

func userLogin(c *gin.Context) {

	failedMessage := "user not registered or password is wrong"
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, failed(err.Error()))
		return
	}

	var realUser models.User
	err = db.Orm.Where(&models.User{Email: user.Email}).Find(&realUser).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, failed(err.Error()))
		return
	}

	if realUser.ID == 0 {
		c.JSON(http.StatusBadRequest, failed(failedMessage))
		return
	}
	m := md5.New()
	m.Write([]byte(user.Password + realUser.Salt))
	formPassword := hex.EncodeToString(m.Sum(nil))
	if formPassword != realUser.Password {
		c.JSON(http.StatusBadRequest, failed(failedMessage))
		return
	}

	session := sessions.Default(c)
	session.Set(auth.UserKey, realUser)
	err = session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, failed(err.Error()))
		return
	}

	c.JSON(http.StatusOK, success())

}

func userLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(auth.UserKey)
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, failed(err.Error()))
		return
	}
	c.JSON(http.StatusOK, success())
}
