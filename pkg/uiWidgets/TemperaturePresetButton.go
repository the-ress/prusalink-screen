package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/domain"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

type TemperaturePresetButton struct {
	*gtk.Button

	printer           *domain.PrinterService
	imageFileName     string
	temperaturePreset *dataModels.TemperaturePreset
	callback          func()
}

func CreateTemperaturePresetButton(
	printer *domain.PrinterService,
	config *config.ScreenConfig,
	imageFileName string,
	temperaturePreset *dataModels.TemperaturePreset,
	callback func(),
) *TemperaturePresetButton {
	presetName := utils.StrEllipsisLen(temperaturePreset.Name, 10)
	base := uiUtils.MustButtonImageUsingFilePath(config, presetName, imageFileName, nil)

	instance := &TemperaturePresetButton{
		Button:            base,
		printer:           printer,
		imageFileName:     imageFileName,
		temperaturePreset: temperaturePreset,
		callback:          callback,
	}
	instance.Button.Connect("clicked", instance.handleClicked)

	return instance
}

func (this *TemperaturePresetButton) handleClicked() {
	logger.Infof("TemperaturePresetButton.handleClicked() - setting temperature to preset %s.", this.temperaturePreset.Name)
	logger.Infof("TemperaturePresetButton.handleClicked() - setting hotend temperature to %.0f.", this.temperaturePreset.Extruder)
	logger.Infof("TemperaturePresetButton.handleClicked() - setting bed temperature to %.0f.", this.temperaturePreset.Bed)

	// Set the bed's temp.
	err := this.printer.SetBedTemperature(this.temperaturePreset.Bed)
	if err != nil {
		logger.LogError("TemperaturePresetButton.handleClicked()", "SetBedTemperature", err)
		return
	}

	// Set the hotend's temp.
	err = this.printer.SetHotendTemperature(this.temperaturePreset.Extruder)
	if err != nil {
		logger.LogError("TemperaturePresetButton.handleClicked()", "SetHotendTemperature", err)
	}

	if this.callback != nil {
		this.callback()
	}
}
