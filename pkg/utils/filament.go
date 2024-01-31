package utils

import (
	// "errors"

	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	// "github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
)

func Extrude(
	client *prusaLinkApis.Client,
	isForward bool,
	parentWindow *gtk.Window,
	flowRatePercentage int,
	length int,
) {
	var action string
	if isForward {
		action = "extrude"
	} else {
		action = "retract"
	}

	if CheckIfHotendTemperatureIsTooLow(client, action, parentWindow) {
		logger.Error("filament.Extrude() - temperature is too low")
		// No need to display an error - CheckIfHotendTemperatureIsTooLow() displays an error to the user
		// if the temperature is too low.
		return
	}

	if err := SetFlowRate(client, flowRatePercentage); err != nil {
		// TODO: display error?
		return
	}

	if err := SendExtrudeRrequest(client, isForward, length); err != nil {
		// TODO: display error?
	}
}

func SetFlowRate(
	client *prusaLinkApis.Client,
	flowRatePercentage int,
) error {
	cmd := &prusaLinkApis.ToolFlowRateRequest{}
	cmd.Factor = flowRatePercentage

	logger.Infof("filament.SetFlowRate() - changing flow rate to %d%%", cmd.Factor)
	if err := cmd.Do(client); err != nil {
		logger.LogError("filament.SetFlowRate()", "Go(ToolFlowRateRequest)", err)
		return err
	}

	return nil
}

func SendExtrudeRrequest(
	client *prusaLinkApis.Client,
	isForward bool,
	length int,
) error {
	cmd := &prusaLinkApis.ToolExtrudeRequest{}
	if isForward {
		cmd.Amount = length
	} else {
		cmd.Amount = -length
	}

	logger.Infof("filament.SendExtrudeRrequest() - sending extrude request with length of: %d", cmd.Amount)
	if err := cmd.Do(client); err != nil {
		logger.LogError("filament.Extrude()", "Do(ToolExtrudeRequest)", err)
		return err
	}

	return nil
}
