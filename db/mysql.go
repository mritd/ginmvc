package db

import (
	"sync"

	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/mritd/ginmvc/utils"
	"github.com/spf13/viper"
)

type orm struct {
	*gorm.DB
}

var Orm orm
var ormOnce sync.Once

func InitMySQL() {

	ormOnce.Do(func() {
		addr := viper.GetString("basic.mysql")
		debug := viper.GetBool("basic.debug")

		db, err := gorm.Open("mysql", addr)
		utils.CheckAndExit(err)

		// Disable table name's pluralization
		db.SingularTable(true)
		db.LogMode(debug)

		Orm = orm{db}
		logrus.Info("mysql init success...")
	})

}
