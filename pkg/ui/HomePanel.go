package ui

import (
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/uiWidgets"
)

type homePanel struct {
	CommonPanel
}

var homePanelInstance *homePanel

func GetHomePanelInstance(
	ui *UI,
) *homePanel {
	if homePanelInstance == nil {
		instance := &homePanel{
			CommonPanel: CreateCommonPanel("HomePanel", ui),
		}
		instance.initialize()
		homePanelInstance = instance
	}

	return homePanelInstance
}

func (this *homePanel) initialize() {
	defer this.Initialize()

	homeXButton := uiWidgets.CreateHomeButton(this.UI.Client, this.UI.Config, "Home X", "home-x.svg", dataModels.XAxis)
	this.Grid().Attach(homeXButton, 2, 1, 1, 1)

	homeYButton := uiWidgets.CreateHomeButton(this.UI.Client, this.UI.Config, "Home Y", "home-y.svg", dataModels.YAxis)
	this.Grid().Attach(homeYButton, 1, 0, 1, 1)

	homeZButton := uiWidgets.CreateHomeButton(this.UI.Client, this.UI.Config, "Home Z", "home-z.svg", dataModels.ZAxis)
	this.Grid().Attach(homeZButton, 1, 1, 1, 1)

	homeAllButton := uiWidgets.CreateHomeAllButton(this.UI.Client, this.UI.Config)
	this.Grid().Attach(homeAllButton, 2, 0, 1, 1)
}
