package config

import "github.com/spf13/viper"

func InitConfig() {
	viper.SetConfigFile("application.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read config file")
	}
}
