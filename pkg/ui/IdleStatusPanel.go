package ui

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/the-ress/prusalink-screen/pkg/domain"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
	"github.com/the-ress/prusalink-screen/pkg/uiWidgets"
)

type idleStatusPanel struct {
	CommonPanel

	nozzleButton *uiWidgets.ToolButton
	bedButton    *uiWidgets.ToolButton
}

var idleStatusPanelInstance *idleStatusPanel

func getIdleStatusPanelInstance(ui *UI) *idleStatusPanel {
	if idleStatusPanelInstance == nil {
		idleStatusPanelInstance = &idleStatusPanel{
			CommonPanel: CreateTopLevelCommonPanel("IdleStatusPanel", ui),
		}
		idleStatusPanelInstance.initialize()
		go idleStatusPanelInstance.consumeStateUpdates(ui.Printer.GetStateUpdates())
	}

	return idleStatusPanelInstance
}

func (this *idleStatusPanel) initialize() {
	logger.TraceEnter("IdleStatusPanel.initialize()")

	defer this.Initialize()

	logger.Info("IdleStatusPanel.initialize() - settings are:")
	if this.UI == nil {
		logger.Error("IdleStatusPanel.initialize() - this.UI is nil")
	}

	menuItems := this.UI.MenuStructure

	menuGrid := uiUtils.MustGrid()
	menuGrid.SetRowHomogeneous(true)
	menuGrid.SetColumnHomogeneous(true)
	this.Grid().Attach(menuGrid, 2, 0, 2, 2)
	this.arrangeMenuItems(menuGrid, menuItems, 2)

	printImage, err := this.UI.ImageLoader.GetImage(uiUtils.PrintSvg)
	if err != nil {
		panic(err)
	}
	printButton := uiUtils.MustButtonImageStyle(printImage, "Print", "color2", this.showFiles)
	this.Grid().Attach(printButton, 2, 2, 2, 1)

	this.showTools()

	logger.TraceLeave("IdleStatusPanel.initialize()")
}

func (this *idleStatusPanel) showFiles() {
	logger.TraceEnter("IdleStatusPanel.showFiles()")

	this.UI.GoToPanel(GetFilesPanelInstance(this.UI))

	logger.TraceLeave("IdleStatusPanel.showFiles()")
}

func (this *idleStatusPanel) showTools() {
	logger.TraceEnter("IdleStatusPanel.showTools()")

	this.nozzleButton = uiWidgets.CreateToolButton(0, this.UI.Printer, this.UI.ImageLoader)
	this.bedButton = uiWidgets.CreateToolButton(-1, this.UI.Printer, this.UI.ImageLoader)

	toolGrid := uiUtils.MustGrid()
	toolGrid.SetRowHomogeneous(true)
	toolGrid.SetColumnHomogeneous(true)
	this.Grid().Attach(toolGrid, 0, 0, 2, 3)
	toolGrid.Attach(this.nozzleButton, 0, 0, 2, 1)
	toolGrid.Attach(this.bedButton, 0, 1, 2, 1)

	logger.TraceLeave("IdleStatusPanel.showTools()")
}

func (this *idleStatusPanel) consumeStateUpdates(ch chan domain.PrinterState) {
	logger.TraceEnter("IdleStatusPanel.consumeStateUpdates()")

	for state := range ch {
		glib.IdleAdd(func() {
			if state.IsConnectedToPrinter {
				this.updateTemperature(state.Temperature)
			}
		})
	}

	logger.TraceLeave("IdleStatusPanel.consumeStateUpdates()")
}

func (this *idleStatusPanel) updateTemperature(temperature dataModels.TemperatureData) {
	this.nozzleButton.SetTemperatures(temperature.Nozzle)
	this.bedButton.SetTemperatures(temperature.Bed)
}
