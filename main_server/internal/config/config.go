package config

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	NetworkConfig NetworkConfig `mapstructure:"network"`
}

type NetworkConfig struct {
	IP   string `mapstructure:"ip"`
	Port string `mapstructure:"port"`
}

func GetConfig(logger *slog.Logger) (*Config, error) {
	logger.Info("read application instance")

	instance := &Config{}

	logger.Info("get path of server config")
	configFileName := os.Getenv("SERVER_CONFIG")

	logger.Info("set config file", slog.String("config", configFileName))
	viper.SetConfigFile(configFileName)

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("config read", slog.String("err", err.Error()))
		return nil, err
	}

	logger.Info("unmarshal config instance")

	if err := viper.Unmarshal(instance); err != nil {
		logger.Error("config unmarshal", slog.String("err", err.Error()))
		return nil, err
	}

	return instance, nil
}
