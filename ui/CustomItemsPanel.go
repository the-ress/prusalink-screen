package ui

import (
	// "github.com/the-ress/prusalink-screen/interfaces"
	// "github.com/the-ress/prusalink-screen/octoprintApis"
	"github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
	// "github.com/the-ress/prusalink-screen/uiWidgets"
)

type customItemsPanel struct {
	CommonPanel
	items []dataModels.MenuItem
}

func CreateCustomItemsPanel(
	ui *UI,
	items []dataModels.MenuItem,
) *customItemsPanel {
	instance := &customItemsPanel{
		CommonPanel: CreateCommonPanel("CustomItemsPanel", ui),
		items:       items,
	}
	instance.initialize()

	return instance
}

func (this *customItemsPanel) initialize() {
	defer this.Initialize()
	this.arrangeMenuItems(this.grid, this.items, 4)
}
