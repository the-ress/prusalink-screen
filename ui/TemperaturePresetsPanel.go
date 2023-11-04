package ui

import (
	// "github.com/Z-Bolt/OctoScreen/interfaces"

	"github.com/Z-Bolt/OctoScreen/octoprintApis/dataModels"

	// "github.com/Z-Bolt/OctoScreen/octoprintApis/dataModels"
	"github.com/Z-Bolt/OctoScreen/uiWidgets"
	// "github.com/Z-Bolt/OctoScreen/utils"
)

type temperaturePresetsPanel struct {
	CommonPanel

	selectHotendStepButton *uiWidgets.SelectToolStepButton
}

var temperaturePresetsPanelInstance *temperaturePresetsPanel

func GetTemperaturePresetsPanelInstance(
	ui *UI,
	selectHotendStepButton *uiWidgets.SelectToolStepButton,
) *temperaturePresetsPanel {
	if temperaturePresetsPanelInstance == nil {
		instance := &temperaturePresetsPanel{
			CommonPanel:            CreateCommonPanel("temperaturePresetsPanel", ui),
			selectHotendStepButton: selectHotendStepButton,
		}
		instance.initialize()
		temperaturePresetsPanelInstance = instance
	}

	return temperaturePresetsPanelInstance
}

func (this *temperaturePresetsPanel) initialize() {
	defer this.Initialize()
	this.createAllOffButton()
	this.createTemperaturePresetButtons()
}

func (this *temperaturePresetsPanel) createAllOffButton() {
	allOffButton := uiWidgets.CreateCoolDownButton(this.UI.Client, this.UI.GoToPreviousPanel)
	this.AddButton(allOffButton)
}

func (this *temperaturePresetsPanel) createTemperaturePresetButtons() {
	// 12 (max) - Back button - All Off button = 10 available slots to display.
	const maxSlots = 10

	// TODO config file
	temperaturePresets := []*dataModels.TemperaturePreset{
		&dataModels.TemperaturePreset{
			Name:     "PLA",
			Extruder: 215,
			Bed:      60,
		},
		&dataModels.TemperaturePreset{
			Name:     "PET",
			Extruder: 230,
			Bed:      85,
		},
		&dataModels.TemperaturePreset{
			Name:     "ASA",
			Extruder: 260,
			Bed:      105,
		},
		&dataModels.TemperaturePreset{
			Name:     "PC",
			Extruder: 275,
			Bed:      110,
		},
		&dataModels.TemperaturePreset{
			Name:     "PVB",
			Extruder: 215,
			Bed:      75,
		},
		&dataModels.TemperaturePreset{
			Name:     "PA",
			Extruder: 275,
			Bed:      90,
		},
		&dataModels.TemperaturePreset{
			Name:     "ABS",
			Extruder: 255,
			Bed:      100,
		},
		&dataModels.TemperaturePreset{
			Name:     "HIPS",
			Extruder: 220,
			Bed:      100,
		},
		&dataModels.TemperaturePreset{
			Name:     "PP",
			Extruder: 254,
			Bed:      100,
		},
		&dataModels.TemperaturePreset{
			Name:     "FLEX",
			Extruder: 240,
			Bed:      50,
		},
	}

	count := 0
	for _, temperaturePreset := range temperaturePresets {
		if count < maxSlots {
			temperaturePresetButton := uiWidgets.CreateTemperaturePresetButton(
				this.UI.Client,
				this.selectHotendStepButton,
				"heat-up.svg",
				temperaturePreset,
				this.UI.GoToPreviousPanel,
			)
			this.AddButton(temperaturePresetButton)
			count++
		}
	}
}
