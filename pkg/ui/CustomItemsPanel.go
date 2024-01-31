package ui

import (
	// "github.com/the-ress/prusalink-screen/pkg/interfaces"
	// "github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	// "github.com/the-ress/prusalink-screen/pkg/uiWidgets"
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
