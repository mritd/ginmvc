package conf

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// since viper does not support concurrent reads, we deserialize to a variable
var Basic Config

type Config struct {
	configPath    string
	Addr          string      `yaml:"addr"`
	Port          int         `yaml:"port"`
	Debug         bool        `yaml:"debug"`
	SessionSecret string      `yaml:"session_secret"`
	LogFile       string      `yaml:"logfile"`
	MySQL         string      `yaml:"mysql"`
	Redis         RedisConfig `yaml:"redis"`
	JWT           JWTConfig   `yaml:"jwt"`
	// if true, we will auto migrate db schema
	AutoMigrate bool `json:"auto_migrate" yaml:"auto_migrate"`
}

type RedisConfig struct {
	DB       int    `yaml:"db"`
	Addr     string `yaml:"addr"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type JWTConfig struct {
	Secret        string `yaml:"secret"`
	SigningMethod string `yaml:"signing_method"`
}

// generate basic example config
func ExampleConfig() Config {
	return Config{
		Addr:          "0.0.0.0",
		Port:          8080,
		Debug:         true,
		SessionSecret: "ARWdeuHoNQjLXTm6rsRLFYMcTvXWtkHD",
		LogFile:       "stdout",
		AutoMigrate:   false,
		MySQL:         "user:password@tcp(test.mysql.com)/dbname?charset=utf8&parseTime=True&loc=Local",
		Redis: RedisConfig{
			Addr: "test.redis.com",
			Port: 6379,
		},
		JWT: JWTConfig{
			Secret:        "aYgDKecXWPn2Jhhs3RPtGCJYPPXxZojr",
			SigningMethod: "HS256",
		},
	}
}

// set config file path
func (cfg *Config) SetConfigPath(configPath string) {
	cfg.configPath = configPath
}

// write config
func (cfg *Config) Write() error {
	if cfg.configPath == "" {
		return errors.New("config path not set")
	}
	out, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(cfg.configPath, out, 0644)
}

// write config to yaml file
func (cfg *Config) WriteTo(filePath string) error {
	if filePath == "" {
		return errors.New("file path is empty")
	}
	cfg.configPath = filePath
	return cfg.Write()
}

// load config
func (cfg *Config) Load() error {
	if cfg.configPath == "" {
		return errors.New("config path not set")
	}
	buf, err := ioutil.ReadFile(cfg.configPath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(buf, cfg)
}

// load config from yaml file
func (cfg *Config) LoadFrom(filePath string) error {
	if filePath == "" {
		return errors.New("file path is empty")
	}
	cfg.configPath = filePath
	return cfg.Load()
}
