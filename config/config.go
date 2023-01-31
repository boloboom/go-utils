package config

import "github.com/spf13/viper"

var config *viper.Viper

func Init(env string) {
	config = viper.New()
	config.SetConfigName(env)
	config.SetConfigType("json")
	config.AddConfigPath("config")
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func GetViper() *viper.Viper {
	return config
}