package config

import "github.com/sirupsen/logrus"

type AppEnv string

const (
	ProductionEnv AppEnv = "production"
	StageEnv      AppEnv = "stage"
	DevelopEnv    AppEnv = "develop"
	LocalEnv      AppEnv = "local"
	TestEnv       AppEnv = "test"
)

type (
	Config struct {
		AppEnv      AppEnv
		LogLevel    logrus.Level
		HTTP        HTTP
		Database    Database
		WorkerCount int
	}

	HTTP struct {
		Port int
	}

	Database struct {
		Postgres Postgres
	}

	Postgres struct {
		Host     string
		Port     int
		Username string
		Password string
		Database string
	}
)
