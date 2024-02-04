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

	moveImage, err := this.UI.ImageLoader.GetImage(uiUtils.MoveSvg)
	if err != nil {
		panic(err)
	}
	moveButton := uiUtils.MustButtonImageStyle(moveImage, "Move", "color1", this.showMove)
	this.Grid().Attach(moveButton, 0, 0, 1, 1)

	filamentSpoolImage, err := this.UI.ImageLoader.GetImage(uiUtils.FilamentSpoolSvg)
	if err != nil {
		panic(err)
	}
	filamentButton := uiUtils.MustButtonImageStyle(filamentSpoolImage, "Filament", "color2", this.showFilament)
	this.Grid().Attach(filamentButton, 1, 0, 1, 1)

	heatUpImage, err := this.UI.ImageLoader.GetImage(uiUtils.HeatUpSvg)
	if err != nil {
		panic(err)
	}
	temperatureButton := uiUtils.MustButtonImageStyle(heatUpImage, "Temperature", "color3", this.showTemperature)
	this.Grid().Attach(temperatureButton, 2, 0, 1, 1)

	networkImage, err := this.UI.ImageLoader.GetImage(uiUtils.NetworkSvg)
	if err != nil {
		panic(err)
	}
	networkButton := uiUtils.MustButtonImageStyle(networkImage, "Network", "color1", this.showNetwork)
	this.Grid().Attach(networkButton, 3, 0, 1, 1)

	infoImage, err := this.UI.ImageLoader.GetImage(uiUtils.InfoSvg)
	if err != nil {
		panic(err)
	}
	systemButton := uiUtils.MustButtonImageStyle(infoImage, "System", "color2", this.showSystem)
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
