package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	HTTP struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"http"`

	Postgres struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Database string `mapstructure:"database"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"postgres"`

	Mongodb struct {
		Host       string `mapstructure:"host"`
		Port       int    `mapstructure:"port"`
		Database   string `mapstructure:"database"`
		Collection string `mapstructure:"collection"`
		User       string `mapstructure:"user"`
		Password   string `mapstructure:"password"`
	} `mapstructure:"mongodb"`
}

func init() {
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot load variables from .env file:\n%s", err.Error())
	}
}

func NewConfig() (cfg Config) {
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("cannot unmarshall configs from file:\n%s", err.Error())
	}
	log.Println("loaded all configurations")
	return cfg
}
