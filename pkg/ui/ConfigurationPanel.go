package ui

import "github.com/the-ress/prusalink-screen/pkg/uiUtils"

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

	networkImage, err := this.UI.ImageLoader.GetImage(uiUtils.NetworkSvg)
	if err != nil {
		panic(err)
	}
	networkButton := uiUtils.MustButtonImageStyle(
		networkImage,
		"Network",
		"color3",
		this.showNetworkPanel,
	)
	this.Grid().Attach(networkButton, 2, 0, 1, 1)

	infoImage, err := this.UI.ImageLoader.GetImage(uiUtils.InfoSvg)
	if err != nil {
		panic(err)
	}

	systemButton := uiUtils.MustButtonImageStyle(
		infoImage,
		"System",
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
