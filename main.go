package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/mritd/ginmvc/db"
	"github.com/mritd/ginmvc/middleware"
	"github.com/mritd/ginmvc/models"

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

		//cache.InitRedis()
		db.InitMySQL()

		// auto migrate db
		if viper.GetBool("basic.auto_migrate") {
			models.AutoMigrate()
		}

		routers.Init()
		middleware.Setup()
		routers.Setup()

		engine := routers.Engine()

		addr := fmt.Sprint(viper.GetString("basic.addr"), ":", viper.GetInt("basic.port"))
		utils.CheckAndExit(engine.Run(addr))

	},
}

func init() {
	// all init func in here
	cobra.OnInitialize(initConfig, initLog)
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

func initLog() {
	debug := viper.GetBool("basic.debug")
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	var logFile io.Writer
	var err error
	logFilePath := viper.GetString("basic.log")
	if strings.ToLower(logFilePath) != "" && strings.ToLower(logFilePath) != "stdout" {
		logFile, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
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
