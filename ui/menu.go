package ui

import (
	"encoding/json"

	"github.com/the-ress/prusalink-screen/interfaces"
	"github.com/the-ress/prusalink-screen/logger"
	"github.com/the-ress/prusalink-screen/octoprintApis"
	"github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
	// "github.com/the-ress/prusalink-screen/uiWidgets"
)

func getPanel(
	ui *UI,
	parentPanel interfaces.IPanel,
	menuItem dataModels.MenuItem,
) interfaces.IPanel {
	switch menuItem.Panel {
	// The standard "top four" panels that are in the idleStatus panel
	case "home":
		return GetHomePanelInstance(ui)

	case "menu":
		fallthrough
	case "custom_items":
		return CreateCustomItemsPanel(ui, menuItem.Items)

	case "filament":
		return GetFilamentPanelInstance(ui)

	case "configuration":
		return GetConfigurationPanelInstance(ui)

	case "files":
		return GetFilesPanelInstance(ui)

	case "temperature":
		return GetTemperaturePanelInstance(ui)

	case "network":
		return GetNetworkPanelInstance(ui)

	case "move":
		return GetMovePanelInstance(ui)

	case "system":
		return GetSystemPanelInstance(ui)

	case "print-menu":
		return GetPrintMenuPanelInstance(ui)

	default:
		logLevel := logger.LogLevel()
		if logLevel == "debug" {
			logger.Fatalf("menu.getPanel() - unknown menuItem.Panel: %q", menuItem.Panel)
		}

		return nil
	}
}

func getDefaultMenuItems(client *octoprintApis.Client) []dataModels.MenuItem {
	defaultMenuItemsForSingleToolhead := `[
		{
			"name": "Home",
			"icon": "home",
			"panel": "home"
		},
		{
			"name": "Actions",
			"icon": "actions",
			"panel": "custom_items",
			"items": [
				{
					"name": "Move",
					"icon": "move",
					"panel": "move"
				},
				{
					"name": "Filament",
					"icon": "filament-spool",
					"panel": "filament"
				},
				{
					"name": "Temperature",
					"icon": "heat-up",
					"panel": "temperature"
				},
				{
					"name": "Network",
					"icon": "network",
					"panel": "network"
				},
				{
					"name": "System",
					"icon": "info",
					"panel": "system"
				}
			]
		},
		{
			"name": "Filament",
			"icon": "filament-spool",
			"panel": "filament"
		},
		{
			"name": "Temperature",
			"icon": "heat-up",
			"panel": "temperature"
		}
	]`

	var menuItems []dataModels.MenuItem
	var err error

	err = json.Unmarshal([]byte(defaultMenuItemsForSingleToolhead), &menuItems)

	if err != nil {
		logger.LogError("menu.getDefaultMenuItems()", "json.Unmarshal()", err)
	}

	return menuItems
}
