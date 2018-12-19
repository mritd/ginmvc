package models

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/mritd/ginmvc/db"
	"github.com/mritd/ginmvc/utils"
)

var migrates []interface{}
var migratesOnce sync.Once

func migrate(obj interface{}) {
	migrates = append(migrates, obj)
}

func AutoMigrate() {
	migratesOnce.Do(func() {
		utils.CheckAndExit(db.Orm.AutoMigrate(migrates...).Error)
		logrus.Info("auto migrate db table success...")
	})
}
