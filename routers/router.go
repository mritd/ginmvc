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

func Engine() *gin.Engine {
	return engine
}

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

// user router setting
func Setup() {
	userRouterOnce.Do(func() {
		for k, f := range routerMap {
			f(engine)
			logrus.Infof("load router [%s] success...", k)
		}
	})
}

func success() gin.H {
	return gin.H{
		"message":   "success",
		"timestamp": time.Now().Unix(),
	}
}

func failed(message string) gin.H {
	return gin.H{
		"message":   message,
		"timestamp": time.Now().Unix(),
	}
}

func data(data interface{}) gin.H {
	return gin.H{
		"message":   "success",
		"data":      data,
		"timestamp": time.Now().Unix(),
	}
}
