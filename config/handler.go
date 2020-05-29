package config

import (
	"git.blp.de/BLP/loggy"
	"github.com/spf13/viper"
)

var _internal *configuration = &configuration{}

func setDefaults() {
	viper.SetDefault("prefix", "+")
}

func Initialize() {
	setDefaults()

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		loggy.Errorf("Could not read configuration (%s).", err.Error())
		panic("Could not read configuration.")
	}

	err := viper.Unmarshal(_internal)

	if err != nil {
		panic(err.Error())
	}
}

func APIToken() string {
	return _internal.Apitoken
}

func Prefix() string {
	return _internal.Prefix
}
