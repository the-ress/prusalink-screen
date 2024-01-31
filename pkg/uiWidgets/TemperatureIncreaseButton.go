package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"

	// "github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

type TemperatureIncreaseButton struct {
	*gtk.Button

	client                      *prusaLinkApis.Client
	temperatureAmountStepButton *TemperatureAmountStepButton
	selectHotendStepButton      *SelectToolStepButton
	isIncrease                  bool
}

func CreateTemperatureIncreaseButton(
	client *prusaLinkApis.Client,
	config *utils.ScreenConfig,
	temperatureAmountStepButton *TemperatureAmountStepButton,
	selectHotendStepButton *SelectToolStepButton,
	isIncrease bool,
) *TemperatureIncreaseButton {
	var base *gtk.Button
	if isIncrease {
		base = utils.MustButtonImageStyle(config, "Increase", "increase.svg", "", nil)
	} else {
		base = utils.MustButtonImageStyle(config, "Decrease", "decrease.svg", "", nil)
	}

	instance := &TemperatureIncreaseButton{
		Button:                      base,
		client:                      client,
		temperatureAmountStepButton: temperatureAmountStepButton,
		selectHotendStepButton:      selectHotendStepButton,
		isIncrease:                  isIncrease,
	}
	instance.Button.Connect("clicked", instance.handleClicked)

	return instance
}

func (this *TemperatureIncreaseButton) handleClicked() {
	value := this.temperatureAmountStepButton.Value()
	tool := this.selectHotendStepButton.Value()
	target, err := utils.GetToolTarget(this.client, tool)
	if err != nil {
		logger.LogError("TemperatureIncreaseButton.handleClicked()", "GetToolTarget()", err)
		return
	}

	if this.isIncrease {
		target += value
	} else {
		target -= value
	}

	if target < 0 {
		target = 0
	}

	// TODO: should the target be checked for a max temp?
	// If so, how to calculate what the max should be?

	logger.Infof("TemperatureIncreaseButton.handleClicked() - setting target temperature for %s to %1.f°C.", tool, target)

	err = utils.SetToolTarget(this.client, tool, target)
	if err != nil {
		logger.LogError("TemperatureIncreaseButton.handleClicked()", "GetToolTarget()", err)
	}
}
