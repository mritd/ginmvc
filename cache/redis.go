package cache

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/mritd/ginmvc/utils"

	"github.com/go-redis/redis"
	"github.com/mritd/ginmvc/conf"
	"github.com/spf13/viper"
)

type redisClient struct {
	*redis.Client
}

var Redis redisClient
var redisOnce sync.Once

func InitRedis() {

	redisOnce.Do(func() {
		var cfg conf.RedisConfig
		utils.CheckAndExit(viper.UnmarshalKey("basic.redis", &cfg))

		Redis = redisClient{
			redis.NewClient(&redis.Options{
				Addr:     fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port),
				Password: cfg.Password,
				DB:       cfg.DB,
			}),
		}

		_, err := Redis.Ping().Result()
		utils.CheckAndExit(err)

		logrus.Info("redis init success...")
	})
}
