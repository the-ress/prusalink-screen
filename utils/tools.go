package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/logger"
	"github.com/the-ress/prusalink-screen/octoprintApis"
	"github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
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

func GetToolTarget(client *octoprintApis.Client, tool string) (float64, error) {
	logger.TraceEnter("Tools.GetToolTarget()")

	fullStateRespone, err := (&octoprintApis.FullStateRequest{
		Exclude: []string{"sd", "state"},
	}).Do(client)

	if err != nil {
		logger.LogError("tools.GetToolTarget()", "Do(StateRequest)", err)
		logger.TraceLeave("Tools.GetToolTarget()")
		return -1, err
	}

	currentTemperatureData, ok := fullStateRespone.Temperature.CurrentTemperatureData[tool]
	if !ok {
		logger.TraceLeave("Tools.GetToolTarget()")
		return -1, fmt.Errorf("unable to find tool %q", tool)
	}

	logger.TraceLeave("Tools.GetToolTarget()")
	return currentTemperatureData.Target, nil
}

func SetToolTarget(client *octoprintApis.Client, tool string, target float64) error {
	logger.TraceEnter("Tools.SetToolTarget()")

	if tool == "bed" {
		cmd := &octoprintApis.BedTargetRequest{Target: target}
		logger.TraceLeave("Tools.SetToolTarget()")
		return cmd.Do(client)
	}

	cmd := &octoprintApis.ToolTargetRequest{Targets: map[string]float64{tool: target}}
	logger.TraceLeave("Tools.SetToolTarget()")
	return cmd.Do(client)
}

func GetCurrentTemperatureData(client *octoprintApis.Client) (map[string]dataModels.TemperatureData, error) {
	logger.TraceEnter("Tools.GetCurrentTemperatureData()")

	temperatureDataResponse, err := (&octoprintApis.TemperatureDataRequest{}).Do(client)
	if err != nil {
		logger.LogError("tools.GetCurrentTemperatureData()", "Do(TemperatureDataRequest)", err)
		logger.TraceLeave("Tools.GetCurrentTemperatureData()")
		return nil, err
	}

	if temperatureDataResponse == nil {
		logger.Error("tools.GetCurrentTemperatureData() - temperatureDataResponse is nil")
		logger.TraceLeave("Tools.GetCurrentTemperatureData()")
		return nil, err
	}

	// Can't test for temperatureDataResponse.TemperatureStateResponse == nil (type mismatch)

	if temperatureDataResponse.TemperatureStateResponse.CurrentTemperatureData == nil {
		logger.Error("tools.GetCurrentTemperatureData() - temperatureDataResponse.TemperatureStateResponse.CurrentTemperatureData is nil")
		logger.TraceLeave("Tools.GetCurrentTemperatureData()")
		return nil, err
	}

	/*
			// Comment out the following to test for multiple hotends:
			currentTemperatureData := make(map[string]dataModels.TemperatureData)
			currentTemperatureData["bed"] = dataModels.TemperatureData{Actual:0.1, Target:0, Offset:0}
		    currentTemperatureData["tool0"] = dataModels.TemperatureData{Actual:10.1, Target:0, Offset:0}
		    currentTemperatureData["tool1"] = dataModels.TemperatureData{Actual:20.1, Target:0, Offset:0}
		    currentTemperatureData["tool2"] = dataModels.TemperatureData{Actual:30.1, Target:0, Offset:0}
		    currentTemperatureData["tool3"] = dataModels.TemperatureData{Actual:40.1, Target:0, Offset:0}
		    currentTemperatureData["tool4"] = dataModels.TemperatureData{Actual:50.1, Target:0, Offset:0}
			return currentTemperatureData, nil
	*/

	logger.TraceLeave("Tools.GetCurrentTemperatureData()")
	return temperatureDataResponse.TemperatureStateResponse.CurrentTemperatureData, nil
}

func CheckIfHotendTemperatureIsTooLow(client *octoprintApis.Client, action string, parentWindow *gtk.Window) bool {
	logger.TraceEnter("Tools.CheckIfHotendTemperatureIsTooLow()")

	currentTemperatureData, err := GetCurrentTemperatureData(client)
	if err != nil {
		logger.LogError("tools.CurrentHotendTemperatureIsTooLow()", "GetCurrentTemperatureData()", err)
		logger.TraceLeave("Tools.CheckIfHotendTemperatureIsTooLow()")
		return true
	}

	temperatureData := currentTemperatureData["tool0"]

	// If the temperature of the hotend is too low, display an error.
	if HotendTemperatureIsTooLow(temperatureData, action, parentWindow) {
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

func GetToolheadFileName() string {
	return "toolhead.svg"
}

func GetExtruderFileName() string {
	return "extruder-typeB.svg"
}

func GetHotendFileName() string {
	return "hotend.svg"
}

func GetNozzleFileName() string {
	return "nozzle.svg"
}

func GetTemperatureDataString(temperatureData dataModels.TemperatureData) string {
	return fmt.Sprintf("%.0f°C / %.0f°C", temperatureData.Actual, temperatureData.Target)
}

// TODO: maybe move HotendTemperatureIsTooLow into a hotend utils file?

const MIN_HOTEND_TEMPERATURE = 150.0

func HotendTemperatureIsTooLow(
	temperatureData dataModels.TemperatureData,
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
