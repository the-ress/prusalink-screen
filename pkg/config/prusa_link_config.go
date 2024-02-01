package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/the-ress/prusalink-screen/pkg/common"
	"github.com/the-ress/prusalink-screen/pkg/logger"

	"gopkg.in/ini.v1"
)

type PrusaLinkConfig struct {
	ApiKey string
}

func ReadPrusaLinkConfig() (*PrusaLinkConfig, error) {
	configFilePath := os.Getenv(common.EnvConfigFilePath)

	if configFilePath == "" {
		return nil, errors.New("PrusaLink config file path is empty")
	}

	logger.Infof("Path to PrusaLink's config file: %q", configFilePath)

	data, err := ini.Load(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read ini file: %q", err)
	}

	cfg := &PrusaLinkConfig{}
	cfg.ApiKey = data.Section("service::local").Key("api_key").String()

	apiKey := os.Getenv(common.EnvPrusaLinkApiKey)
	if apiKey != "" {
		cfg.ApiKey = apiKey
	}

	return cfg, nil
}
