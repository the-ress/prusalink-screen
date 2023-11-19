package ui

import (
	// "time"

	// "github.com/Z-Bolt/OctoScreen/interfaces"
	"github.com/Z-Bolt/OctoScreen/logger"
	"github.com/Z-Bolt/OctoScreen/uiWidgets"
	"github.com/Z-Bolt/OctoScreen/utils"
)

type systemPanel struct {
	CommonPanel

	// First row
	octoPrintInfoBox        *uiWidgets.OctoPrintInfoBox
	octoScreenInfoBox       *uiWidgets.OctoScreenInfoBox
	octoScreenPluginInfoBox *uiWidgets.OctoScreenPluginInfoBox

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
		instance.preShowCallback = instance.refreshSystemInformationInfoBox
		systemPanelInstance = instance
	}

	return systemPanelInstance
}

func (this *systemPanel) initialize() {
	logger.TraceEnter("SystemPanel.initialize()")

	defer this.Initialize()

	// First row
	logoWidth := this.Scaled(52)
	this.octoPrintInfoBox = uiWidgets.CreateOctoPrintInfoBox(this.UI.Client, logoWidth)
	this.Grid().Attach(this.octoPrintInfoBox, 0, 0, 1, 1)

	this.octoScreenInfoBox = uiWidgets.CreateOctoScreenInfoBox(this.UI.Client, utils.OctoScreenVersion)
	this.Grid().Attach(this.octoScreenInfoBox, 1, 0, 2, 1)

	this.octoScreenPluginInfoBox = uiWidgets.CreateOctoScreenPluginInfoBox(this.UI.Client, this.UI.OctoPrintPluginIsAvailable)
	this.Grid().Attach(this.octoScreenPluginInfoBox, 3, 0, 1, 1)

	// Second row
	this.systemInformationInfoBox = uiWidgets.CreateSystemInformationInfoBox(this.UI.window, this.UI.scaleFactor)
	this.Grid().Attach(this.systemInformationInfoBox, 0, 1, 4, 1)

	logger.TraceLeave("SystemPanel.initialize()")
}

func (this *systemPanel) refreshSystemInformationInfoBox() {
	this.systemInformationInfoBox.Refresh()
}
