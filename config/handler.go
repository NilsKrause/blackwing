package config

import (
	"github.com/spf13/viper"
)

var _internal *configuration = &configuration{}

// gets called from Initialize to set the default values in the configuration
func setDefaults() {
	viper.SetDefault("prefix", "+")
}

// Initialize reads the config file and sets defaults
func Initialize() {
	setDefaults()

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic("Could not read configuration.")
	}

	err := viper.Unmarshal(_internal)

	if err != nil {
		panic(err.Error())
	}
}

// APIToken returns the, from the configuration set APIToken
func APIToken() string {
	return _internal.Apitoken
}

// Prefix sets the bot prefix
func Prefix() string {
	return _internal.Prefix
}
