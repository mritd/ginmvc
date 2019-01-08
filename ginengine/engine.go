package ginengine

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/mritd/ginmvc/conf"
	"github.com/sirupsen/logrus"
)

var Engine *gin.Engine
var engineOnce sync.Once

// init gin router engine
func Init() {
	engineOnce.Do(func() {
		if conf.Basic.Debug {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}

		Engine = gin.New()
		logrus.Info("create gin engine success...")
	})

}
