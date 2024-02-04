package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type FilamentExtrudeButton struct {
	*gtk.Button

	parentWindow              *gtk.Window
	client                    *prusaLinkApis.Client
	amountToExtrudeStepButton *AmountToExtrudeStepButton
	flowRateStepButton        *FlowRateStepButton // The flow rate step button is optional.
	isForward                 bool
}

func CreateFilamentExtrudeButton(
	parentWindow *gtk.Window,
	client *prusaLinkApis.Client,
	imageLoader *uiUtils.ImageLoader,
	amountToExtrudeStepButton *AmountToExtrudeStepButton,
	flowRateStepButton *FlowRateStepButton, // The flow rate step button is optional.
	isForward bool,
) *FilamentExtrudeButton {
	var base *gtk.Button
	if isForward {
		image, err := imageLoader.GetImage(uiUtils.ExtruderExtrudeSvg)
		if err != nil {
			panic(err)
		}
		base = uiUtils.MustButtonImageStyle(image, "Extrude", "", nil)
	} else {
		image, err := imageLoader.GetImage(uiUtils.ExtruderRetractSvg)
		if err != nil {
			panic(err)
		}
		base = uiUtils.MustButtonImageStyle(image, "Retract", "", nil)
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

	uiUtils.Extrude(
		this.client,
		this.isForward,
		this.parentWindow,
		flowRatePercentage,
		length,
	)
}
