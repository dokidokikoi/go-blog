package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Config() *viper.Viper {
	return config
}

func SetConfig(filename string) {
	config = viper.New()
	config.SetConfigFile(filename)
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func init() {
	SetConfig("./internal/conf/application.yaml")
}

func GetSpecConfig(key string) *viper.Viper {
	return config.Sub(key)
}
