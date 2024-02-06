package utils

import (
	"fmt"
	"os"

	"github.com/the-ress/prusalink-screen/pkg/common"
	"github.com/the-ress/prusalink-screen/pkg/logger"
)

const MISSING_ENV_TOKEN = ">>MISSING<<"
const INVALID_ENV_TOKEN = "!!!INVALID!!!"

func DumpSystemInformation() {
	logger.Info("System Information...")
	logger.Infof("PrusaLink Screen version: %q", common.AppVersion)
	logger.Infof("PrusaLink Screen build date: %q", common.BuildDate)
	// More system stats to come...
	logger.Info("")
}

func DumpEnvironmentVariables() {
	logger.Info("Environment variables...")

	// Required environment variables
	logger.Info("Required environment variables:")
	dumpEnvironmentVariable(common.EnvPrusaLinkHost)

	// TODO: revisit this!
	// 1. remove OCTOPRINT_APIKEY from option settings
	// 2. make the octoprint config path required
	// 3. update code... use only one path to get the api key octoprint)
	// 4. update code... make octoprint config path required
	// 5. update code... read api key from octoprint config
	// 6. dump api key (obfuscated though)
	// 7. update docs
	// 8. make sure what's dumped to the log is correct, for both when present and when missing.
	dumpObfuscatedEnvironmentVariable(common.EnvPrusaLinkApiKey)

	dumpEnvironmentVariable(common.EnvStylePath)
	logger.Info("")

	// Optional environment variables
	logger.Info("Optional environment variables:")
	dumpEnvironmentVariable(common.EnvConfigFilePath)
	dumpEnvironmentVariable(common.EnvPrusaLinkExecutablePath)
	dumpEnvironmentVariable(common.EnvPrusaLinkUser)
	dumpEnvironmentVariable(common.EnvLogFilePath)
	dumpEnvironmentVariable(common.EnvLogLevel)

	dumpEnvironmentVariable(common.EnvResolution)
	// EnvResolution is optional.  If not set, the window size will
	// default to the values defined in globalVars.go.

	dumpEnvironmentVariable(common.EnvDisplayCursor)
	logger.Info("")
}

func dumpEnvironmentVariable(key string) {
	value := os.Getenv(key)
	if value == "" {
		value = MISSING_ENV_TOKEN
	}

	logger.Infof("key: %q, value: %q", key, value)
}

func dumpObfuscatedEnvironmentVariable(key string) {
	value := os.Getenv(key)
	if value == "" {
		value = MISSING_ENV_TOKEN
	} else {
		value = GetObfuscatedValue(value)
	}

	logger.Infof("key: %q, value: %q", key, value)
}

func GetObfuscatedValue(value string) string {
	length := len(value)

	obfuscatedValue := ""
	if length < 6 {
		obfuscatedValue = INVALID_ENV_TOKEN
	} else {
		if value == MISSING_ENV_TOKEN {
			return value
		} else {
			obfuscatedValue = fmt.Sprintf("%c%c%c---%c%c%c",
				value[0],
				value[1],
				value[2],
				value[length-3],
				value[length-2],
				value[length-1],
			)
		}
	}

	return obfuscatedValue
}
