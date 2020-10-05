package main

import (
	"fmt"
	"go-rest-api/api"
	"go-rest-api/config"
	"log"

	"github.com/spf13/viper"
)

var (
	conf config.Configuration
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		if error, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(fmt.Sprintf("Config File Not found,\n%s", error))
		} else {
			log.Fatal(fmt.Sprintf("Error when reading config file,\n%s", err))
		}
	}

	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Errorf("Unmarshall Error, %s", err)
	}

	app := &api.App{}
	app.InitAndServe(&conf)
}
