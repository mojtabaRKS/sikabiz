package config

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Load() (*Config, error) {

	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.AllowEmptyEnv(true)
	viper.SetDefault("APP_ENV", LocalEnv)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("reading config: %w", err)
		}
	}

	logLvl, err := logrus.ParseLevel(viper.GetString("LOG_LEVEL"))
	if err != nil {
		return nil, fmt.Errorf("parsing LOG_LEVEL: %w", err)
	}

	return &Config{
		AppEnv:      AppEnv(viper.GetString("APP_ENV")),
		LogLevel:    logLvl,
		WorkerCount: viper.GetInt("WORKER_COUNT"),
		HTTP: HTTP{
			Port: viper.GetInt("HTTP_PORT"),
		},
		Database: Database{
			Postgres: Postgres{
				Host:     viper.GetString("POSTGRES_HOST"),
				Port:     viper.GetInt("POSTGRES_PORT"),
				Username: viper.GetString("POSTGRES_USER"),
				Password: viper.GetString("POSTGRES_PASSWORD"),
				Database: viper.GetString("POSTGRES_DB"),
			},
		},
	}, nil
}
