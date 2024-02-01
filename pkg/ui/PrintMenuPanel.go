package ui

import "github.com/the-ress/prusalink-screen/pkg/uiUtils"

// "github.com/the-ress/prusalink-screen/pkg/interfaces"

type printMenuPanel struct {
	CommonPanel
}

var printMenuPanelInstance *printMenuPanel

func GetPrintMenuPanelInstance(
	ui *UI,
) *printMenuPanel {
	if printMenuPanelInstance == nil {
		instance := &printMenuPanel{
			CommonPanel: CreateCommonPanel("PrintMenuPanel", ui),
		}
		instance.initialize()
		printMenuPanelInstance = instance
	}

	return printMenuPanelInstance
}

func (this *printMenuPanel) initialize() {
	defer this.Initialize()

	moveButton := uiUtils.MustButtonImageStyle(this.UI.Config, "Move", "move.svg", "color1", this.showMove)
	this.Grid().Attach(moveButton, 0, 0, 1, 1)

	filamentButton := uiUtils.MustButtonImageStyle(this.UI.Config, "Filament", "filament-spool.svg", "color2", this.showFilament)
	this.Grid().Attach(filamentButton, 1, 0, 1, 1)

	temperatureButton := uiUtils.MustButtonImageStyle(this.UI.Config, "Temperature", "heat-up.svg", "color3", this.showTemperature)
	this.Grid().Attach(temperatureButton, 2, 0, 1, 1)

	networkButton := uiUtils.MustButtonImageStyle(this.UI.Config, "Network", "network.svg", "color1", this.showNetwork)
	this.Grid().Attach(networkButton, 3, 0, 1, 1)

	systemButton := uiUtils.MustButtonImageStyle(this.UI.Config, "System", "info.svg", "color2", this.showSystem)
	this.Grid().Attach(systemButton, 0, 1, 1, 1)
}

func (this *printMenuPanel) showMove() {
	this.UI.GoToPanel(GetMovePanelInstance(this.UI))
}

func (this *printMenuPanel) showFilament() {
	this.UI.GoToPanel(GetFilamentPanelInstance(this.UI))
}

func (this *printMenuPanel) showTemperature() {
	this.UI.GoToPanel(GetTemperaturePanelInstance(this.UI))
}

func (this *printMenuPanel) showNetwork() {
	this.UI.GoToPanel(GetNetworkPanelInstance(this.UI))
}

func (this *printMenuPanel) showSystem() {
	this.UI.GoToPanel(GetSystemPanelInstance(this.UI))
}
