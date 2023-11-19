package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/logger"
	"github.com/the-ress/prusalink-screen/octoprintApis"
	"github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
	"github.com/the-ress/prusalink-screen/utils"
)

type HomeAllButton struct {
	*gtk.Button

	client *octoprintApis.Client
}

func CreateHomeAllButton(
	client *octoprintApis.Client,
) *HomeAllButton {
	base := utils.MustButtonImageStyle("Home All", "home.svg", "", nil)

	instance := &HomeAllButton{
		Button: base,
		client: client,
	}
	instance.Button.Connect("clicked", instance.handleClicked)

	return instance
}

func (this *HomeAllButton) handleClicked() {
	logger.Infof("Homing the print head")

	// Version A:
	axes := []dataModels.Axis{
		dataModels.XAxis,
		dataModels.YAxis,
		dataModels.ZAxis,
	}
	cmd := &octoprintApis.PrintHeadHomeRequest{Axes: axes}
	err := cmd.Do(this.client)
	if err != nil {
		logger.LogError("HomeAllButton.handleClicked()", "Do(PrintHeadHomeRequest)", err)
	}

	/*
		// If there are issues with version A, there's also version B:
		cmd := &octoprint.CommandRequest{}
		cmd.Commands = []string{
			"G28 Z",
			"G28 X",
			"G28 Y",
		}

		if err := cmd.Do(m.UI.Client); err != nil {
			logger.LogError("HomeAllButton.handleClicked()", "Do(CommandRequest)", err)
		}
	*/
}
