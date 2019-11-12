package db

import (
	"strings"
	"sync"

	"github.com/mritd/ginmvc/conf"
	"github.com/mritd/ginmvc/utils"

	"github.com/jmoiron/sqlx"

	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	*sqlx.DB
}

var MySQL mysql
var mysqlOnce sync.Once

// init mysql sqlx
func InitMySQL() {
	mysqlOnce.Do(func() {
		db, err := sqlx.Connect("mysql", conf.Basic.MySQL)
		utils.CheckAndExit(err)

		db.MapperFunc(strings.ToLower)
		MySQL = mysql{db}
		logrus.Info("mysql init success...")
	})
}
