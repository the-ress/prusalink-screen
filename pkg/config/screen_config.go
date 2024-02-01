package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/gdk"
	"github.com/the-ress/prusalink-screen/pkg/common"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

type ScreenConfig struct {
	// Required configs
	PrusaLinkApiKey         string
	PrusaLinkHost           string
	PrusaLinkExecutablePath string
	PrusaLinkUser           string
	CssStyleFilePath        string

	// Optional configs
	LogFilePath   string
	LogLevel      string
	WindowSize    screenConfigWindowSize
	DisplayCursor bool
}

type screenConfigWindowSize struct {
	Width  int
	Height int
}

func ReadConfig() (*ScreenConfig, error) {
	prusaLinkConfig, err := ReadPrusaLinkConfig()

	apiKey := ""
	if err != nil {
		logger.Warnf("Failed to load PrusaLink config: %q", err)
	} else {
		apiKey = prusaLinkConfig.ApiKey
	}

	config := &ScreenConfig{
		// Required configs
		PrusaLinkApiKey:  apiKey,
		PrusaLinkHost:    common.DefaultServerHost,
		CssStyleFilePath: "", // default to "" for now, but this must be set in the environment variables

		// Optional configs
		LogFilePath: "",
		LogLevel:    "",
		WindowSize:  readWindowSize(),
		// Resolution: screenConfigWindowSize{
		// 	Width:  -1,
		// 	Height: -1,
		// },
		DisplayCursor: false,
	}

	readEnvironmentValues(config)

	url := strings.ToLower(config.PrusaLinkHost)
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		logger.Warn("PRUSALINK_HOST requires the transport protocol ('http://' or 'https://') but is missing. 'http://' is being added.")
		config.PrusaLinkHost = fmt.Sprintf("http://%s", config.PrusaLinkHost)
	}

	err = validateConfig(config)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func readEnvironmentValues(config *ScreenConfig) {
	host := os.Getenv(common.EnvPrusaLinkHost)
	if host != "" {
		config.PrusaLinkHost = host
	}

	apiKey := os.Getenv(common.EnvPrusaLinkApiKey)
	if apiKey != "" {
		config.PrusaLinkApiKey = apiKey
	}

	executablePath := os.Getenv(common.EnvPrusaLinkExecutablePath)
	if executablePath != "" {
		config.PrusaLinkExecutablePath = executablePath
	}

	user := os.Getenv(common.EnvPrusaLinkUser)
	if user != "" {
		config.PrusaLinkUser = user
	}

	cssStyleFilePath := os.Getenv(common.EnvStylePath)
	if cssStyleFilePath != "" {
		config.CssStyleFilePath = cssStyleFilePath
	}

	logFilePath := os.Getenv(common.EnvLogFilePath)
	if logFilePath != "" {
		config.LogFilePath = logFilePath
	}

	logLevel := os.Getenv(common.EnvLogLevel)
	if logLevel != "" {
		config.LogLevel = logLevel
	}

	displayCursor := strings.ToLower(os.Getenv(common.EnvDisplayCursor))
	if displayCursor == "true" {
		config.DisplayCursor = true
	} else {
		config.DisplayCursor = false
	}
}

func getScreenSize() (*screenConfigWindowSize, error) {
	display, err := gdk.DisplayGetDefault()
	if err != nil {
		return nil, fmt.Errorf("Error getting the default display: %q", err)
	}
	monitor, err := display.GetPrimaryMonitor()
	if err != nil {
		return nil, fmt.Errorf("Error getting the primary monitor: %q", err)
	}

	geometry := monitor.GetGeometry()

	return &screenConfigWindowSize{
		Width:  geometry.GetWidth(),
		Height: geometry.GetHeight(),
	}, nil
}

func readWindowSize() screenConfigWindowSize {
	sizeString := os.Getenv(common.EnvResolution)

	size, err := parseWindowSizeString(sizeString)

	if size != nil {
		return *size
	} else if err != nil {
		logger.Errorf("Failed to parse window size: %q", err)
	}

	logger.Info("Detecting screen size")
	size, err = getScreenSize()

	if size != nil {
		return *size
	} else if err != nil {
		logger.Errorf("Failed to detect screen size: %q", err)
	}

	logger.Infof("Using default window size: %dx%d", common.DefaultWindowWidth, common.DefaultWindowHeight)

	return screenConfigWindowSize{
		Width:  common.DefaultWindowWidth,
		Height: common.DefaultWindowHeight,
	}
}

func parseWindowSizeString(resolutionString string) (*screenConfigWindowSize, error) {
	if resolutionString == "" {
		logger.Infof("%s is empty", common.EnvResolution)
		return nil, nil
	}

	parts := strings.SplitN(resolutionString, "x", 2)

	if len(parts) != 2 {
		return nil, fmt.Errorf("%s is malformed: %q", common.EnvResolution, resolutionString)
	}

	width, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("Failed to parse window width (%s): %q", parts[0], err)
	} else if width < common.MinimumWindowWidth || width > common.MaximumWindowWidth {
		return nil, fmt.Errorf(
			"Window width (%d) is out of range (%d, %d)",
			width,
			common.MinimumWindowWidth,
			common.MaximumWindowWidth,
		)
	}

	height, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("Failed to parse window height (%s): %q", parts[1], err)
	} else if height < common.MinimumWindowHeight || height > common.MaximumWindowHeight {
		return nil, fmt.Errorf(
			"Window height (%d) is out of range (%d, %d)",
			height,
			common.MinimumWindowHeight,
			common.MaximumWindowHeight,
		)
	}

	return &screenConfigWindowSize{
		Width:  width,
		Height: height,
	}, nil
}

func validateConfig(config *ScreenConfig) error {
	if config.PrusaLinkApiKey == "" {
		return errors.New("Required setting is not set: PrusaLinkApiKey")
	}

	if config.PrusaLinkHost == "" {
		return errors.New("Required setting is not set: PrusaLinkHost")
	}

	if config.CssStyleFilePath == "" {
		return errors.New("Required setting is not set: CssStyleFilePath")
	}

	return nil
}

func (this *ScreenConfig) Dump() {
	logger.Info("Dumping config:")

	// Required configs
	logger.Infof("%-16s: %q", "PrusaLinkApiKey", utils.GetObfuscatedValue(this.PrusaLinkApiKey))
	logger.Infof("%-16s: %q", "PrusaLinkHost", this.PrusaLinkHost)
	logger.Infof("%-16s: %q", "PrusaLinkExecutablePath", this.PrusaLinkExecutablePath)
	logger.Infof("%-16s: %q", "PrusaLinkUser", this.PrusaLinkUser)
	logger.Infof("%-16s: %q", "CssStyleFilePath", this.CssStyleFilePath)

	// Optional configs
	logger.Infof("%-16s: %q", "LogFilePath", this.LogFilePath)
	logger.Infof("%-16s: %q", "LogLevel", this.LogLevel)
	logger.Infof("%-16s: %d", "WindowSize.Width", this.WindowSize.Width)
	logger.Infof("%-16s: %d", "WindowSize.Height", this.WindowSize.Height)
	logger.Infof("%-16s: %t", "DisplayCursor", this.DisplayCursor)

	logger.Info("")
}
