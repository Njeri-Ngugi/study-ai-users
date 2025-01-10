package config

import (
	"github.com/Njeri-Ngugi/toolbox/config"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func FromEnv() (cfg *config.GlobalConfig, err error) {
	fromFileToEnv()

	cfg = &config.GlobalConfig{}

	err = envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func fromFileToEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Debug("No config files found to load to env. Defaulting to environment.")
	} else {
		logrus.Debug("Config files found loading to env.")
	}
}
