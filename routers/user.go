package routers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/mritd/ginmvc/auth"
	"github.com/mritd/ginmvc/db"
	"github.com/mritd/ginmvc/middleware"
	"github.com/mritd/ginmvc/models"
	"github.com/mritd/ginmvc/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	register("user", func(router *gin.Engine) {
		v1UserGroup := router.Group("api/v1/user")
		v1UserGroup.POST("/register", userRegister)
		v1UserGroup.POST("/login", userLogin)
		v1UserGroup.POST("/logout", middleware.SessionAuth, userLogout)
	})
}

// @Tags User API
// @Summary user register
// @Description user register
// @Accept json
// @Produce json
// @Param user body models.User true "user info"
// @Success 200 {object} models.CommonResp "{"message":"success","timestamp":1570887849}"
// @Failure 400 {object} models.CommonResp "{"message":"user email [mritd@linux.com] already register","timestamp":1570887884}"
// @Failure 500 {object} models.CommonResp "{"message":"invalid connection","timestamp":1570887977}"
// @Router /api/v1/user/register [post]
func userRegister(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, failed(err.Error()))
		return
	}

	// check user email if already exist
	var count int
	err = db.MySQL.Get(&count, "SELECT COUNT(1) FROM t_user WHERE email = ?", user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, failed(err.Error()))
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, failed("user email [%s] already register", user.Email))
		return
	}

	// create user

	salt := utils.RandString(auth.SaltLength)
	m := md5.New()
	m.Write([]byte(user.Password.String + salt))
	password := hex.EncodeToString(m.Sum(nil))

	addUserSQL := "INSERT INTO t_user (name, email, mobile, password, `lock`, salt, create_time,update_time,login_time) VALUES (?,?,?,?,?,?,?,?,?)"
	_, err = db.MySQL.Exec(addUserSQL, user.Name, user.Email, user.Mobile, password, 0, salt, time.Now().Unix(), 0, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, failed(err.Error()))
		return
	}
	c.JSON(http.StatusOK, success())
}

// @Tags User API
// @Summary user login
// @Description user login
// @Accept json
// @Produce json
// @Param user body models.User true "user info"
// @Success 200 {object} models.CommonResp "{"message":"success","timestamp":1570889247}"
// @Failure 400 {object} models.CommonResp "{"message":"user not registered or password is wrong","timestamp":1570889272}"
// @Failure 500 {object} models.CommonResp "{"message":"invalid connection","timestamp":1570887977}"
// @Router /api/v1/user/login [post]
func userLogin(c *gin.Context) {

	failedMessage := "user not registered or password is wrong"

	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, failed(err.Error()))
		return
	}

	var realUser models.User
	querySQL := "SELECT * FROM t_user WHERE email = ?"
	err = db.MySQL.Get(&realUser, querySQL, user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, failed(failedMessage))
			return
		} else {
			c.JSON(http.StatusInternalServerError, failed(err.Error()))
			return
		}
	}

	m := md5.New()
	m.Write([]byte(user.Password.String + realUser.Salt.String))
	formPassword := hex.EncodeToString(m.Sum(nil))
	if formPassword != realUser.Password.String {
		c.JSON(http.StatusBadRequest, failed(failedMessage))
		return
	}

	loginSQL := "UPDATE t_user SET login_time = ? WHERE email = ?"
	_, err = db.MySQL.Exec(loginSQL, time.Now().Unix(), user.Email)
	if err != nil {
		c.JSON(500, failed(err.Error()))
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

// @Tags User API
// @Summary user logout
// @Description user logout
// @Produce json
// @Success 200 {object} models.CommonResp "{"message":"success","timestamp":1570889247}"
// @Failure 500 {object} models.CommonResp "{"message":"error message","timestamp":1570887977}"
// @Router /api/v1/user/logout [post]
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
