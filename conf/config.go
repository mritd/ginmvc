package conf

import (
	"net"
)

type Config struct {
	Addr          net.IP      `json:"addr" yaml:"addr" mapstructure:"addr"`
	Port          int         `json:"port" yaml:"port" mapstructure:"port"`
	Debug         bool        `json:"debug" yaml:"debug" mapstructure:"debug"`
	SessionSecret string      `json:"session_secret" yaml:"session_secret" mapstructure:"session_secret"`
	MySQL         string      `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	Redis         RedisConfig `json:"redis" yaml:"redis" mapstructure:"redis"`
	AutoMigrate   bool        `json:"auto_migrate" yaml:"auto_migrate" mapstructure:"auto_migrate"`
}

type RedisConfig struct {
	DB       int    `json:"db" yaml:"db" mapstructure:"db"`
	Addr     string `json:"addr" yaml:"addr" mapstructure:"addr"`
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
}

func ExampleConfig() *Config {
	return &Config{
		Addr:          net.ParseIP("0.0.0.0"),
		Port:          8080,
		Debug:         true,
		SessionSecret: "ARWdeuHoNQjLXTm6rsRLFYMcTvXWtkHD",
		AutoMigrate:   false,
		MySQL:         "user:password@tcp(test.mysql.com)/dbname?charset=utf8&parseTime=True&loc=Local",
		Redis: RedisConfig{
			Addr: "test.redis.com",
			Port: 6379,
		},
	}
}
