package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/mritd/ginmvc/models"

	"github.com/mritd/ginmvc/db"
	"github.com/mritd/ginmvc/middleware"
	"github.com/sirupsen/logrus"

	"github.com/mritd/ginmvc/routers"

	"github.com/mritd/ginmvc/utils"

	"github.com/mritd/ginmvc/conf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "ginmvc",
	Short: "Gin mvc template",
	Long: `
Gin mvc template.`,
	Run: func(cmd *cobra.Command, args []string) {

		// load config
		conf.Load()
		// init framework log
		initLog()
		// current we don't need cache
		// if one day needs it, we can delete the comment
		//cache.InitRedis()
		// init mysql(gorm)
		db.InitMySQL()
		// migrate db schema
		models.AutoMigrate()
		// init gin router engine
		routers.Init()
		// load middleware
		middleware.Setup()
		// add gin router
		routers.Setup()

		// run gin http server
		engine := routers.Engine()
		addr := fmt.Sprint(conf.Basic.Addr, ":", conf.Basic.Port)
		logrus.Infof("server listen at %s", addr)
		utils.CheckAndExit(engine.Run(addr))

	},
}

func init() {
	// load config file
	cobra.OnInitialize(initConfig)
	// cmd config flag
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "ginmvc.yaml", "config file (default is ginmvc.yaml)")
}

func initConfig() {

	viper.SetConfigFile(cfgFile)

	if _, err := os.Stat(cfgFile); err != nil {
		_, err := os.Create(cfgFile)
		utils.CheckAndExit(err)
		viper.Set("basic", conf.ExampleConfig())
		utils.CheckAndExit(viper.WriteConfig())
	}

	viper.AutomaticEnv()
	utils.CheckAndExit(viper.ReadInConfig())
}

// init log config
func initLog() {
	if conf.Basic.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	var logFile io.Writer
	var err error
	if strings.ToLower(conf.Basic.LogFile) != "" && strings.ToLower(conf.Basic.LogFile) != "stdout" {
		logFile, err = os.OpenFile(conf.Basic.LogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			logrus.Fatal(err)
		}
	} else {
		logFile = os.Stdout
	}

	logrus.SetOutput(logFile)
	logrus.Infof("GOMAXPROCS: %d", runtime.NumCPU())
}

func main() {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)
	utils.CheckAndExit(rootCmd.Execute())
}
