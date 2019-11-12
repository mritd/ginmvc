package auth

import (
	"sync"

	"github.com/mritd/ginmvc/conf"

	sqlxadapter "github.com/memwey/casbin-sqlx-adapter"

	"github.com/casbin/casbin"
	"github.com/mritd/ginmvc/utils"
	"github.com/sirupsen/logrus"
)

const CasbinRBACModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
`

var Enforcer *casbin.Enforcer
var casbinOnce sync.Once

func InitCasbin() {
	casbinOnce.Do(func() {
		// refs => https://github.com/memwey/casbin-sqlx-adapter
		adapter := sqlxadapter.NewAdapterFromOptions(&sqlxadapter.AdapterOptions{
			DriverName:     "mysql",
			DataSourceName: conf.Basic.MySQL,
			TableName:      "casbin_rule",
			//DB:             nil,
		})
		Enforcer = casbin.NewEnforcer(casbin.NewModel(CasbinRBACModel), adapter)
		utils.CheckAndExit(Enforcer.LoadPolicy())
		logrus.Info("casbin init success...")
	})

}
