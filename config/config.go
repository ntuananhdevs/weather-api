package config

import (
	"log"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile("../.env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Không thể đọc file config: %v", err)
	}
}

func GetAPIKey() string {
	return viper.GetString("API_KEY")
}
