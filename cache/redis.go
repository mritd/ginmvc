package cache

import (
	"fmt"
	"sync"

	"github.com/mritd/ginmvc/conf"

	"github.com/sirupsen/logrus"

	"github.com/mritd/ginmvc/utils"

	"github.com/go-redis/redis"
)

type redisClient struct {
	*redis.Client
}

var Redis redisClient
var redisOnce sync.Once

func InitRedis() {

	redisOnce.Do(func() {

		Redis = redisClient{
			redis.NewClient(&redis.Options{
				Addr:     fmt.Sprintf("%s:%d", conf.Basic.Redis.Addr, conf.Basic.Redis.Port),
				Password: conf.Basic.Redis.Password,
				DB:       conf.Basic.Redis.DB,
			}),
		}

		_, err := Redis.Ping().Result()
		utils.CheckAndExit(err)

		logrus.Info("redis init success...")
	})
}
