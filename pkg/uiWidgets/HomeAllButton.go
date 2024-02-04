package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type HomeAllButton struct {
	*gtk.Button

	client *prusaLinkApis.Client
}

func CreateHomeAllButton(
	client *prusaLinkApis.Client,
	imageLoader *uiUtils.ImageLoader,
) *HomeAllButton {
	image, err := imageLoader.GetImage(uiUtils.HomeSvg)
	if err != nil {
		panic(err)
	}
	base := uiUtils.MustButtonImageStyle(image, "Home All", "", nil)

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
	cmd := &prusaLinkApis.PrintHeadHomeRequest{Axes: axes}
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
