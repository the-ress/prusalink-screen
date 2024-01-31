package ui

import (
	"github.com/the-ress/prusalink-screen/utils"
)

type configurationPanel struct {
	CommonPanel
}

var configurationPanelInstance *configurationPanel

func GetConfigurationPanelInstance(
	ui *UI,
) *configurationPanel {
	if configurationPanelInstance == nil {
		instance := &configurationPanel{
			CommonPanel: CreateCommonPanel("ConfigurationPanel", ui),
		}
		instance.initialize()
		configurationPanelInstance = instance
	}

	return configurationPanelInstance
}

func (this *configurationPanel) initialize() {
	defer this.Initialize()

	networkButton := utils.MustButtonImageStyle(
		this.UI.Config,
		"Network",
		"network.svg",
		"color3",
		this.showNetworkPanel,
	)
	this.Grid().Attach(networkButton, 2, 0, 1, 1)

	systemButton := utils.MustButtonImageStyle(
		this.UI.Config,
		"System",
		"info.svg",
		"color4",
		this.showSystemPanel,
	)
	this.Grid().Attach(systemButton, 3, 0, 1, 1)
}

func (this *configurationPanel) showNetworkPanel() {
	this.UI.GoToPanel(GetNetworkPanelInstance(this.UI))
}

func (this *configurationPanel) showSystemPanel() {
	this.UI.GoToPanel(GetSystemPanelInstance(this.UI))
}
