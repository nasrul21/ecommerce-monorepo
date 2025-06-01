package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	GRPC struct {
		Port int `mapstructure:"PORT"`
	} `mapstructure:"GRPC"`
	DB struct {
		Host     string `mapstructure:"HOST"`
		Port     int    `mapstructure:"PORT"`
		User     string `mapstructure:"USER"`
		Password string `mapstructure:"PASSWORD"`
		Name     string `mapstructure:"NAME"`
		SSLMode  string `mapstructure:"SSL_MODE"`
	} `mapstructure:"DB"`
	Auth struct {
		Token struct {
			Expired time.Duration `mapstructure:"EXPIRED"`
			Secret  string        `mapstructure:"SECRET"`
		} `mapstructure:"TOKEN"`
	} `mapstructure:"AUTH"`
}

func LoadConfig() *Config {
	cfg := Config{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't find the file .env: ", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &cfg
}
