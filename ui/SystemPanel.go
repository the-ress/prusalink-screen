package ui

import (
	// "time"

	// "github.com/the-ress/prusalink-screen/interfaces"

	"github.com/the-ress/prusalink-screen/uiWidgets"
	"github.com/the-ress/prusalink-screen/utils"
)

type systemPanel struct {
	CommonPanel

	// First row
	prusaLinkInfoBox       *uiWidgets.PrusaLinkInfoBox
	prusaConnectInfoBox    *uiWidgets.OctoScreenPluginInfoBox
	prusaLinkScreenInfoBox *uiWidgets.OctoScreenInfoBox

	// Second row
	systemInformationInfoBox *uiWidgets.SystemInformationInfoBox
}

var systemPanelInstance *systemPanel = nil

func GetSystemPanelInstance(
	ui *UI,
) *systemPanel {
	if systemPanelInstance == nil {
		instance := &systemPanel{
			CommonPanel: CreateCommonPanel("SystemPanel", ui),
		}
		instance.initialize()
		instance.preShowCallback = instance.refreshData
		systemPanelInstance = instance
	}

	return systemPanelInstance
}

func (this *systemPanel) initialize() {
	defer this.Initialize()

	// First row
	logoWidth := this.Scaled(52)
	this.prusaLinkInfoBox = uiWidgets.CreatePrusaLinkInfoBox(this.UI.Client, this.UI.Printer, logoWidth)
	this.Grid().Attach(this.prusaLinkInfoBox, 0, 0, 1, 1)

	this.prusaConnectInfoBox = uiWidgets.CreateOctoScreenPluginInfoBox(this.UI.Client, this.UI.OctoPrintPluginIsAvailable)
	this.Grid().Attach(this.prusaConnectInfoBox, 1, 0, 2, 1)

	this.prusaLinkScreenInfoBox = uiWidgets.CreateOctoScreenInfoBox(this.UI.Client, utils.OctoScreenVersion)
	this.Grid().Attach(this.prusaLinkScreenInfoBox, 3, 0, 1, 1)

	// Second row
	this.systemInformationInfoBox = uiWidgets.CreateSystemInformationInfoBox(this.UI.window, this.UI.scaleFactor)
	this.Grid().Attach(this.systemInformationInfoBox, 0, 1, 4, 1)
}

func (this *systemPanel) refreshData() {
	this.systemInformationInfoBox.Refresh()
}
