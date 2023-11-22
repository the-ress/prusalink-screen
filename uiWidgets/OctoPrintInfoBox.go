package uiWidgets

import (
	"fmt"

	"github.com/the-ress/prusalink-screen/domain"
	"github.com/the-ress/prusalink-screen/logger"
	"github.com/the-ress/prusalink-screen/octoprintApis"

	// "github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
	"github.com/the-ress/prusalink-screen/utils"
)

type OctoPrintInfoBox struct {
	*SystemInfoBox
}

func CreateOctoPrintInfoBox(
	client *octoprintApis.Client,
	printer *domain.PrinterService,
	logoWidth int,
) *OctoPrintInfoBox {
	logger.TraceEnter("OctoPrintInfoBox.CreateOctoPrintInfoBox()")

	logoHeight := int(float64(logoWidth) * 1.25)
	logoImage := utils.MustImageFromFileWithSize("logos/logo-octoprint.png", logoWidth, logoHeight)

	server := "Unknown?"
	apiVersion := "Unknown?"

	if printer.GetState().IsConnectedToPrusaLink {
		// Only call if we're connected to PrusaLink
		versionResponse, err := (&octoprintApis.VersionRequest{}).Do(client)
		if err != nil {
			logger.LogError("OctoPrintInfoBox.CreateOctoPrintInfoBox()", "VersionRequest.Do()", err)
		} else if versionResponse != nil {
			server = versionResponse.Server
			apiVersion = versionResponse.API
		}
	}

	base := CreateSystemInfoBox(
		client,
		logoImage,
		"OctoPrint",
		server,
		fmt.Sprintf("(API   %s)", apiVersion), // Use 3 spaces here... 1 space doesn't have enough kerning.
	)

	instance := &OctoPrintInfoBox{
		SystemInfoBox: base,
	}

	logger.TraceLeave("OctoPrintInfoBox.CreateOctoPrintInfoBox()")

	return instance
}
