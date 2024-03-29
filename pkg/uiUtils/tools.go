package uiUtils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
)

func GetDisplayNameForTool(toolName string) string {
	// Since this is such a hack, lets add some bounds checking
	if toolName == "" {
		logger.Error("Tools..GetDisplayNameForTool() - toolName is empty")
		return ""
	}

	lowerCaseName := strings.ToLower(toolName)
	if strings.LastIndex(lowerCaseName, "tool") != 0 {
		logger.Errorf("Tools.GetDisplayNameForTool() - toolName is invalid, value passed in was: %q", toolName)
		return ""
	}

	if len(toolName) != 5 {
		logger.Errorf("Tools.GetDisplayNameForTool() - toolName is invalid, value passed in was: %q", toolName)
		return ""
	}

	toolIndexAsInt, _ := strconv.Atoi(string(toolName[4]))
	displayName := toolName[0:4]
	displayName = displayName + strconv.Itoa(toolIndexAsInt+1)

	return displayName
}

func GetToolTarget(client *prusaLinkApis.Client, tool string) (float64, error) {

	statusResponse, err := (&prusaLinkApis.StatusRequest{}).Do(client)

	if err != nil {
		logger.LogError("tools.GetToolTarget()", "Do(StatusRequest)", err)
		return -1, err
	}

	if tool == "bed" {
		return statusResponse.Printer.TargetBed, nil
	} else {
		return statusResponse.Printer.TargetNozzle, nil
	}
}

func SetToolTarget(client *prusaLinkApis.Client, tool string, target float64) error {
	logger.TraceEnter("Tools.SetToolTarget()")

	if tool == "bed" {
		cmd := &prusaLinkApis.BedTargetRequest{Target: target}
		logger.TraceLeave("Tools.SetToolTarget()")
		return cmd.Do(client)
	}

	cmd := &prusaLinkApis.ToolTargetRequest{Targets: map[string]float64{tool: target}}
	logger.TraceLeave("Tools.SetToolTarget()")
	return cmd.Do(client)
}

func GetNozzleTemperatureData(client *prusaLinkApis.Client) (*dataModels.ToolTemperatureData, error) {
	temperatureDataResponse, err := (&prusaLinkApis.StatusRequest{}).Do(client)
	if err != nil {
		logger.LogError("tools.GetCurrentTemperatureData()", "Do(StatusRequest)", err)
		return nil, err
	}

	if temperatureDataResponse == nil {
		logger.Error("tools.GetCurrentTemperatureData() - temperatureDataResponse is nil")
		return nil, err
	}

	return &dataModels.ToolTemperatureData{
		Actual: temperatureDataResponse.Printer.TempNozzle,
		Target: temperatureDataResponse.Printer.TargetNozzle,
	}, nil
}

func CheckIfHotendTemperatureIsTooLow(client *prusaLinkApis.Client, action string, parentWindow *gtk.Window) bool {
	logger.TraceEnter("Tools.CheckIfHotendTemperatureIsTooLow()")

	temperatureData, err := GetNozzleTemperatureData(client)
	if err != nil {
		logger.LogError("tools.CurrentHotendTemperatureIsTooLow()", "GetNozzleTemperatureData()", err)
		logger.TraceLeave("Tools.CheckIfHotendTemperatureIsTooLow()")
		return true
	}

	// If the temperature of the hotend is too low, display an error.
	if HotendTemperatureIsTooLow(*temperatureData, action, parentWindow) {
		errorMessage := fmt.Sprintf(
			"The temperature of the hotend is too low to %s.\n(the current temperature is only %.0f°C)\n\nPlease increase the temperature and try again.",
			action,
			temperatureData.Actual,
		)
		ErrorMessageDialogBox(parentWindow, errorMessage)

		logger.TraceLeave("Tools.CheckIfHotendTemperatureIsTooLow()")
		return true
	}

	logger.TraceLeave("Tools.CheckIfHotendTemperatureIsTooLow()")
	return false
}

func GetTemperatureDataString(temperatureData dataModels.ToolTemperatureData) string {
	return fmt.Sprintf("%.0f°C / %.0f°C", temperatureData.Actual, temperatureData.Target)
}

// TODO: maybe move HotendTemperatureIsTooLow into a hotend utils file?

const MIN_HOTEND_TEMPERATURE = 150.0

func HotendTemperatureIsTooLow(
	temperatureData dataModels.ToolTemperatureData,
	action string,
	parentWindow *gtk.Window,
) bool {
	targetTemperature := temperatureData.Target
	logger.Infof("tools.HotendTemperatureIsTooLow() - targetTemperature is %.2f", targetTemperature)

	actualTemperature := temperatureData.Actual
	logger.Infof("tools.HotendTemperatureIsTooLow() - actualTemperature is %.2f", actualTemperature)

	if targetTemperature <= MIN_HOTEND_TEMPERATURE || actualTemperature <= MIN_HOTEND_TEMPERATURE {
		return true
	}

	return false
}
