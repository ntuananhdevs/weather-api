package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)	
	} else {
		log.Println("Config file loaded")
	}
}

func GetApiKey() string {
	return viper.GetString("API_KEY")
}
func main () {
	LoadConfig()
	fmt.Println(GetApiKey())
}
