package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/logger"
	"github.com/the-ress/prusalink-screen/octoprintApis"
	"github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
	"github.com/the-ress/prusalink-screen/utils"
)

type HomeButton struct {
	*gtk.Button

	client *octoprintApis.Client
	axes   []dataModels.Axis
}

func CreateHomeButton(
	client *octoprintApis.Client,
	buttonLabel string,
	imageFileName string,
	axes ...dataModels.Axis,
) *HomeButton {
	base := utils.MustButtonImageStyle(buttonLabel, imageFileName, "", nil)

	instance := &HomeButton{
		Button: base,
		client: client,
		axes:   axes,
	}
	instance.Button.Connect("clicked", instance.handleClicked)

	return instance
}

func (this *HomeButton) handleClicked() {
	cmd := &octoprintApis.PrintHeadHomeRequest{Axes: this.axes}
	logger.Infof("Homing the print head in %s axes", this.axes)
	err := cmd.Do(this.client)
	if err != nil {
		logger.LogError("HomeButton.handleClicked()", "Do(PrintHeadHomeRequest)", err)
	}
}
