package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() *Config {
	viper.AddConfigPath("./")
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

type Config struct {
	// mysql db
	MysqlHost     string `mapstructure:"MYSQL_HOST"`
	MysqlPort     int    `mapstructure:"MYSQL_PORT"`
	MysqlUser     string `mapstructure:"MYSQL_USER"`
	MysqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	MysqlDb       string `mapstructure:"MYSQL_DB"`
	// auth
	AuthSecret             string `mapstructure:"AUTH_SECRET"`
	AuthTokenExpireSeconds int    `mapstructure:"AUTH_TOKEN_EXPIRE_SECONDS"`
	AuthTokenIssuer        string `mapstructure:"AUTH_TOKEN_ISSUER"`
}
