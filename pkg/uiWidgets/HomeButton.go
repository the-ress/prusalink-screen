package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

type HomeButton struct {
	*gtk.Button

	client *prusaLinkApis.Client
	axes   []dataModels.Axis
}

func CreateHomeButton(
	client *prusaLinkApis.Client,
	config *utils.ScreenConfig,
	buttonLabel string,
	imageFileName string,
	axes ...dataModels.Axis,
) *HomeButton {
	base := utils.MustButtonImageStyle(config, buttonLabel, imageFileName, "", nil)

	instance := &HomeButton{
		Button: base,
		client: client,
		axes:   axes,
	}
	instance.Button.Connect("clicked", instance.handleClicked)

	return instance
}

func (this *HomeButton) handleClicked() {
	cmd := &prusaLinkApis.PrintHeadHomeRequest{Axes: this.axes}
	logger.Infof("Homing the print head in %s axes", this.axes)
	err := cmd.Do(this.client)
	if err != nil {
		logger.LogError("HomeButton.handleClicked()", "Do(PrintHeadHomeRequest)", err)
	}
}
