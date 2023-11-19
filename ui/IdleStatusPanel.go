package ui

import (
	// "encoding/json"
	// "fmt"
	// "os"
	// "strconv"
	// "sync"
	// "time"

	// "github.com/Z-Bolt/OctoScreen/octoprintApis"
	"github.com/Z-Bolt/OctoScreen/logger"
	"github.com/Z-Bolt/OctoScreen/octoprintApis/dataModels"
	"github.com/Z-Bolt/OctoScreen/uiWidgets"
	"github.com/Z-Bolt/OctoScreen/utils"
	"github.com/gotk3/gotk3/glib"
)

type idleStatusPanel struct {
	CommonPanel

	tool0Button *uiWidgets.ToolButton
	bedButton   *uiWidgets.ToolButton
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

func GoToIdleStatusPanel(ui *UI) {
	instance := getIdleStatusPanelInstance(ui)
	ui.GoToPanel(instance)
}

func (this *idleStatusPanel) initialize() {
	logger.TraceEnter("IdleStatusPanel.initialize()")

	defer this.Initialize()

	logger.Info("IdleStatusPanel.initialize() - settings are:")
	if this.UI == nil {
		logger.Error("IdleStatusPanel.initialize() - this.UI is nil")
	}

	var menuItems []dataModels.MenuItem
	if this.UI.MenuStructure == nil || len(this.UI.MenuStructure) < 1 {
		logger.Info("IdleStatusPanel.initialize() - Loading default menu")
		this.UI.MenuStructure = getDefaultMenuItems(this.UI.Client)
	} else {
		logger.Info("IdleStatusPanel.initialize() - Loading octo menu")
	}

	menuItems = this.UI.MenuStructure

	menuGrid := utils.MustGrid()
	menuGrid.SetRowHomogeneous(true)
	menuGrid.SetColumnHomogeneous(true)
	this.Grid().Attach(menuGrid, 2, 0, 2, 2)
	this.arrangeMenuItems(menuGrid, menuItems, 2)

	printButton := utils.MustButtonImageStyle("Print", "print.svg", "color2", this.showFiles)
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

	this.tool0Button = uiWidgets.CreateToolButton(0, this.UI.Printer)
	this.bedButton = uiWidgets.CreateToolButton(-1, this.UI.Printer)

	toolGrid := utils.MustGrid()
	toolGrid.SetRowHomogeneous(true)
	toolGrid.SetColumnHomogeneous(true)
	this.Grid().Attach(toolGrid, 0, 0, 2, 3)
	toolGrid.Attach(this.tool0Button, 0, 0, 2, 1)
	toolGrid.Attach(this.bedButton, 0, 1, 2, 1)

	logger.TraceLeave("IdleStatusPanel.showTools()")
}

func (this *idleStatusPanel) consumeStateUpdates(ch chan *dataModels.FullStateResponse) {
	logger.TraceEnter("IdleStatusPanel.consumeStateUpdates()")

	for fullStateResponse := range ch {
		glib.IdleAdd(func() {
			this.updateTemperature(fullStateResponse)
		})
	}

	logger.TraceLeave("IdleStatusPanel.consumeStateUpdates()")
}

func (this *idleStatusPanel) updateTemperature(fullStateResponse *dataModels.FullStateResponse) {
	logger.TraceEnter("IdleStatusPanel.updateTemperature()")

	octoPrintResponseManager := GetOctoPrintResponseManagerInstance(this.UI)
	if octoPrintResponseManager.IsConnected() != true {
		// If not connected, do nothing and leave.
		logger.TraceLeave("IdleStatusPanel.updateTemperature() (not connected)")
		return
	}

	for tool, currentTemperatureData := range fullStateResponse.Temperature.CurrentTemperatureData {
		switch tool {
		case "bed":
			logger.Debug("Updating the UI's bed temp")
			this.bedButton.SetTemperatures(currentTemperatureData)

		case "tool0":
			logger.Debug("Updating the UI's tool0 temp")
			this.tool0Button.SetTemperatures(currentTemperatureData)

		default:
			logger.Errorf("IdleStatusPanel.updateTemperature() - GetOctoPrintResponseManagerInstance() returned an unknown tool: %q", tool)
		}
	}

	logger.TraceLeave("IdleStatusPanel.updateTemperature()")
}
