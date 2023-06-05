package configs

import (
	"go-folder-sample/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitialiseConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("configs/")

	if err := viper.ReadInConfig(); err != nil {
		logger.Logger.Error("fatal error config file")
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			logger.Logger.Fatal("config file not found: ", zap.Error(err))
		} else {
			// Config file was found but another error was produced
			logger.Logger.Fatal("config file was found but another error was produced: ", zap.Error(err))
		}
	}

	viper.SetConfigName("constants")
	viper.SetConfigType("json")
	viper.AddConfigPath("configs/")
	if err := viper.MergeInConfig(); err != nil {
		logger.Logger.Fatal("constants file  error was : ", zap.Error(err))
	}

}
