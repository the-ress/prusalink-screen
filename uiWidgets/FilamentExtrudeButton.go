package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/octoprintApis"

	// "github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
	"github.com/the-ress/prusalink-screen/utils"
)

type FilamentExtrudeButton struct {
	*gtk.Button

	parentWindow              *gtk.Window
	client                    *octoprintApis.Client
	amountToExtrudeStepButton *AmountToExtrudeStepButton
	flowRateStepButton        *FlowRateStepButton // The flow rate step button is optional.
	isForward                 bool
}

func CreateFilamentExtrudeButton(
	parentWindow *gtk.Window,
	client *octoprintApis.Client,
	amountToExtrudeStepButton *AmountToExtrudeStepButton,
	flowRateStepButton *FlowRateStepButton, // The flow rate step button is optional.
	isForward bool,
) *FilamentExtrudeButton {
	var base *gtk.Button
	if isForward {
		base = utils.MustButtonImageStyle("Extrude", "extruder-extrude.svg", "", nil)
	} else {
		base = utils.MustButtonImageStyle("Retract", "extruder-retract.svg", "", nil)
	}

	instance := &FilamentExtrudeButton{
		Button:                    base,
		parentWindow:              parentWindow,
		client:                    client,
		amountToExtrudeStepButton: amountToExtrudeStepButton,
		flowRateStepButton:        flowRateStepButton,
		isForward:                 isForward,
	}
	instance.Button.Connect("clicked", instance.handleClicked)

	return instance
}

func (this *FilamentExtrudeButton) handleClicked() {
	this.sendExtrudeCommand(this.amountToExtrudeStepButton.Value())
}

func (this *FilamentExtrudeButton) sendExtrudeCommand(length int) {
	flowRatePercentage := 100
	if this.flowRateStepButton != nil {
		flowRatePercentage = this.flowRateStepButton.Value()
	}

	utils.Extrude(
		this.client,
		this.isForward,
		this.parentWindow,
		flowRatePercentage,
		length,
	)
}
