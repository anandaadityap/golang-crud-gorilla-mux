package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Config struct {
	PORT          string
	DB_CONNECTION string
}

var ENV Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal(err)
	}

	log.Println("Load server successfully")
}
