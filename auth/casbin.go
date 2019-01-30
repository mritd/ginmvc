package auth

import (
	"sync"

	gormadapter "github.com/casbin/gorm-adapter"

	"github.com/casbin/casbin"
	"github.com/mritd/ginmvc/conf"
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
		adapter := gormadapter.NewAdapter("mysql", conf.Basic.MySQL, true)
		Enforcer = casbin.NewEnforcer(casbin.NewModel(CasbinRBACModel), adapter)
		utils.CheckAndExit(Enforcer.LoadPolicy())
		logrus.Info("casbin init success...")
	})

}
