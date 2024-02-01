package common

// Required environment variables
const (
	EnvStylePath       = "PRUSALINK_SCREEN_STYLE_PATH"
	EnvPrusaLinkHost   = "PRUSALINK_HOST"
	EnvPrusaLinkApiKey = "PRUSALINK_APIKEY"
)

// Optional (but good to have) environment variables
const (
	EnvLogLevel                = "PRUSALINK_SCREEN_LOG_LEVEL"
	EnvLogFilePath             = "PRUSALINK_SCREEN_LOG_FILE_PATH"
	EnvResolution              = "PRUSALINK_SCREEN_RESOLUTION"
	EnvConfigFilePath          = "PRUSALINK_CONFIG_FILE"
	EnvPrusaLinkExecutablePath = "PRUSALINK_EXECUTABLE_PATH"
	EnvPrusaLinkUser           = "PRUSALINK_USER"
	EnvDisplayCursor           = "DISPLAY_CURSOR"
)
