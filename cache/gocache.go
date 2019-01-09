package cache

import (
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var MemCache *cache.Cache
var goCacheOnce sync.Once

func InitMemCache() {
	goCacheOnce.Do(func() {
		MemCache = cache.New(cache.NoExpiration, 10*time.Minute)
		logrus.Info("memory cache init success...")
	})
}
