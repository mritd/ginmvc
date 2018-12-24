package routers

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/mritd/ginmvc/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var engine *gin.Engine
var routerOnce, userRouterOnce sync.Once
var routerMap map[string]func(router *gin.Engine)

// init gin router engine
func Init() {

	routerOnce.Do(func() {
		if viper.GetBool("basic.debug") {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}

		engine = gin.New()
		logrus.Info("create gin engine success...")
	})

}

// return gin router engine instance
func Engine() *gin.Engine {
	return engine
}

// register new router with key name
// key name is used to eliminate duplicate routes
// key name not case sensitive
func register(key string, method func(router *gin.Engine)) {
	if routerMap == nil {
		routerMap = make(map[string]func(router *gin.Engine), 50)
	}
	if routerMap[key] != nil {
		utils.CheckAndExit(errors.New(fmt.Sprintf("method key [%s] already exist!\n", key)))
	} else {
		routerMap[key] = method
	}
}

// framework init
func Setup() {
	userRouterOnce.Do(func() {
		for k, f := range routerMap {
			f(engine)
			logrus.Infof("load router [%s] success...", k)
		}
	})
}

// for the fast return success result
func success() gin.H {
	return gin.H{
		"message":   "success",
		"timestamp": time.Now().Unix(),
	}
}

// for the fast return failed result
func failed(message string) gin.H {
	return gin.H{
		"message":   message,
		"timestamp": time.Now().Unix(),
	}
}

// for the fast return result with custom data
func data(data interface{}) gin.H {
	return gin.H{
		"message":   "success",
		"data":      data,
		"timestamp": time.Now().Unix(),
	}
}
