package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/the-ress/prusalink-screen/logger"

	"gopkg.in/ini.v1"
)

var (
	homeOctoPi = "/home/jo/"
)

type OctoPrintConfig struct {
	ApiKey         string
	Host           string
	ExecutablePath string
	User           string
}

func ReadOctoPrintConfig() *OctoPrintConfig {
	logger.TraceEnter("OctoPrintConfig.ReadOctoPrintConfig()")

	configFilePath := os.Getenv(EnvConfigFilePath)
	if configFilePath == "" {
		configFilePath = findOctoPrintConfigFilePath()
	}

	if configFilePath == "" {
		logger.Info("OctoPrintConfig.ReadOctoPrintConfig() - configFilePath is empty")
		logger.TraceLeave("OctoPrintConfig.ReadOctoPrintConfig()")
		return &OctoPrintConfig{}
	}

	logger.Infof("Path to OctoPrint's config file: %q", configFilePath)

	data, err := ini.Load(configFilePath)
	if err != nil {
		panic(fmt.Sprintf("OctoPrintConfig.ReadOctoPrintConfig() - ini.Load() returned an error: %q", err))
	}

	cfg := &OctoPrintConfig{}
	cfg.ApiKey = data.Section("service::local").Key("api_key").String()

	logger.TraceLeave("OctoPrintConfig.ReadOctoPrintConfig()")
	return cfg
}

func findOctoPrintConfigFilePath() string {
	logger.TraceEnter("OctoPrintConfig.FindOctoPrintConfigFilePath()")

	filePath := filepath.Join(homeOctoPi, "prusa_printer_settings.ini")
	if _, err := os.Stat(filePath); err == nil {
		logger.Info("OctoPrintConfig.FindOctoPrintConfigFilePath() - doFindOctoPrintConfigFilePath() found a file")
		logger.TraceLeave("OctoPrintConfig.FindOctoPrintConfigFilePath(), returning the file")
		return filePath
	}

	logger.TraceLeave("OctoPrintConfig.FindOctoPrintConfigFilePath(), returning an empty string")
	return ""
}

func (this *OctoPrintConfig) OverrideConfigsWithEnvironmentValues() {
	logger.TraceEnter("OctoPrintConfig.OverrideConfigsWithEnvironmentValues()")

	apiKey := os.Getenv(EnvOctoPrintApiKey)
	if apiKey != "" {
		this.ApiKey = apiKey
	}

	host := os.Getenv(EnvOctoPrintHost)
	if host != "" {
		this.Host = host
	}

	executablePath := os.Getenv(EnvPrusaLinkExecutablePath)
	if executablePath != "" {
		this.ExecutablePath = executablePath
	}

	user := os.Getenv(EnvPrusaLinkUser)
	if user != "" {
		this.User = user
	}

	logger.TraceLeave("OctoPrintConfig.OverrideConfigsWithEnvironmentValues()")
}

func (this *OctoPrintConfig) UpdateValues() {
	logger.TraceEnter("OctoPrintConfig.UpdateValues()")

	if this.Host == "" {
		logger.Infof("Server host is empty, defaulting to the default value (%s)", DefaultServerHost)
		this.Host = DefaultServerHost
	}

	url := strings.ToLower(this.Host)
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		logger.Warn("WARNING!  OCTOPRINT_HOST requires the transport protocol ('http://' or 'https://') but is missing.  'http://' is being added to Host.")
		this.Host = fmt.Sprintf("http://%s", this.Host)
	}

	logger.TraceLeave("OctoPrintConfig.UpdateValues()")
}

func (this *OctoPrintConfig) MissingRequiredConfigName() string {
	logger.TraceEnter("OctoPrintConfig.MissingRequiredConfigName()")

	if this.ApiKey == "" {
		return "ApiKey"
	}

	if this.Host == "" || this.Host == "http://" {
		return "Host"
	}

	logger.TraceLeave("OctoPrintConfig.MissingRequiredConfigName()")

	return ""
}

func (this *OctoPrintConfig) DumpConfigs() {
	// Don't add TraceEnter/TraceLeave to this function.

	logger.Infof("%-16s: %q", "ApiKey", GetObfuscatedValue(this.ApiKey))
	logger.Infof("%-16s: %q", "Host", this.Host)
}
