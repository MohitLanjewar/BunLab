package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Configurations struct {
	Env               string
	Prefix            string
	CustomRespDomains []string
	Server            struct {
		Port string
	}
}

func GetConfigurations() *Configurations {

	viper.SetConfigName("config")

	viper.AddConfigPath("./config/data/")
	if os.Getenv("GO_ENV") == "local" {
		viper.AddConfigPath("./config/data/")
	} else {
		viper.AddConfigPath("/etc/config/")
	}

	viper.AutomaticEnv()
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	var configuration Configurations

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	//fmt.Println(configuration)
	return &configuration
}
