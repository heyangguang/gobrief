package cmd

import (
	"fmt"
	"github.com/spf13/viper"
)

func viperInit(config string) {
	viper.SetConfigFile(config)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
