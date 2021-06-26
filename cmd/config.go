package cmd

import (
	"fmt"
	"github.com/spf13/viper"
)

type TotalConfig struct {
	LogPath, LogLevel, DBConnect string
	Port                         int
}

var AllConfig TotalConfig

func viperInit(config string) {
	viper.SetConfigFile(config)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	err = viper.Unmarshal(&AllConfig)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
