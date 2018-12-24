package conf

import (
	"github.com/mritd/ginmvc/utils"
	"github.com/spf13/viper"
)

var Basic Config

type Config struct {
	Addr          string      `json:"addr" yaml:"addr" mapstructure:"addr"`
	Port          int         `json:"port" yaml:"port" mapstructure:"port"`
	Debug         bool        `json:"debug" yaml:"debug" mapstructure:"debug"`
	SessionSecret string      `json:"session_secret" yaml:"session_secret" mapstructure:"session_secret"`
	LogPath       string      `json:"log_path" yaml:"log_path" mapstructure:"log_path"`
	MySQL         string      `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	Redis         RedisConfig `json:"redis" yaml:"redis" mapstructure:"redis"`
	// if true, we will auto migrate db schema
	AutoMigrate bool `json:"auto_migrate" yaml:"auto_migrate" mapstructure:"auto_migrate"`
}

type RedisConfig struct {
	DB       int    `json:"db" yaml:"db" mapstructure:"db"`
	Addr     string `json:"addr" yaml:"addr" mapstructure:"addr"`
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
}

// generate basic example config
func ExampleConfig() *Config {
	return &Config{
		Addr:          "0.0.0.0",
		Port:          8080,
		Debug:         true,
		SessionSecret: "ARWdeuHoNQjLXTm6rsRLFYMcTvXWtkHD",
		LogPath:       "stdout",
		AutoMigrate:   false,
		MySQL:         "user:password@tcp(test.mysql.com)/dbname?charset=utf8&parseTime=True&loc=Local",
		Redis: RedisConfig{
			Addr: "test.redis.com",
			Port: 6379,
		},
	}
}

// load config
func Load() {
	utils.CheckAndExit(viper.UnmarshalKey("basic", &Basic))
}
