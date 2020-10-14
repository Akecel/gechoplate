package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// InitConfig set all configuration for the project
func InitConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err.Error()))
	}
}
