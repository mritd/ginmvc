package db

import (
	"sync"

	"github.com/mritd/ginmvc/conf"

	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/mritd/ginmvc/utils"
)

type orm struct {
	*gorm.DB
}

var Orm orm
var ormOnce sync.Once

// init mysql gorm
func InitMySQL() {
	ormOnce.Do(func() {
		db, err := gorm.Open("mysql", conf.Basic.MySQL)
		utils.CheckAndExit(err)

		// disable table name's pluralization
		db.SingularTable(true)
		db.LogMode(conf.Basic.Debug)
		// set logger to logrus
		db.SetLogger(NewGormrus())

		Orm = orm{db}
		logrus.Info("mysql init success...")
	})
}
