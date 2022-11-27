package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func loadConfig(envFileLoc string) *Config {
	viper.AddConfigPath(envFileLoc)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
func LoadConfig() *Config {
	return loadConfig("./")
}

const rootDir = "go_web_be/"

func LoadConfigForTest() *Config {

	pwd, _ := os.Getwd()
	index := strings.Index(pwd, rootDir)
	if index < 0 {
		panic(fmt.Sprintf("could not found env file for test at: %s", pwd))
	}
	envFileLoc := pwd[0 : index+len(rootDir)]

	return loadConfig(envFileLoc)
}

type Auth struct {
	Secret             string `mapstructure:"AUTH_SECRET"`
	TokenExpireSeconds int    `mapstructure:"AUTH_TOKEN_EXPIRE_SECONDS"`
	TokenIssuer        string `mapstructure:"AUTH_TOKEN_ISSUER"`
}
type Config struct {
	Mysql struct {
		Host     string `mapstructure:"MYSQL_HOST"`
		Port     int    `mapstructure:"MYSQL_PORT"`
		User     string `mapstructure:"MYSQL_USER"`
		Password string `mapstructure:"MYSQL_PASSWORD"`
		Db       string `mapstructure:"MYSQL_DB"`
	} `mapstructure:",squash"`

	Auth struct {
		Secret             string `mapstructure:"AUTH_SECRET"`
		TokenExpireSeconds int    `mapstructure:"AUTH_TOKEN_EXPIRE_SECONDS"`
		TokenIssuer        string `mapstructure:"AUTH_TOKEN_ISSUER"`
	} `mapstructure:",squash"`

	Redis struct {
		Host     string `mapstructure:"REDIS_HOST"`
		Port     int    `mapstructure:"REDIS_PORT"`
		Password string `mapstructure:"REDIS_PASSWORD"`
	} `mapstructure:",squash"`
}
