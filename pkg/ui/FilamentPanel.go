package ui

import (
	// "fmt"
	// "strings"
	// "time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	// "github.com/the-ress/prusalink-screen/pkg/interfaces"
	"github.com/the-ress/prusalink-screen/pkg/domain"
	"github.com/the-ress/prusalink-screen/pkg/uiWidgets"
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

type filamentPanel struct {
	CommonPanel

	// First row
	filamentExtrudeButton     *uiWidgets.FilamentExtrudeButton
	flowRateStepButton        *uiWidgets.FlowRateStepButton
	amountToExtrudeStepButton *uiWidgets.AmountToExtrudeStepButton
	filamentRetractButton     *uiWidgets.FilamentExtrudeButton

	// Second row
	// filamentLoadButton   *uiWidgets.FilamentLoadButton
	temperatureStatusBox *uiWidgets.TemperatureStatusBox
	// filamentUnloadButton *uiWidgets.FilamentLoadButton

	// Third row
	temperatureButton *gtk.Button
}

var filamentPanelInstance *filamentPanel

func GetFilamentPanelInstance(
	ui *UI,
) *filamentPanel {
	if filamentPanelInstance == nil {
		filamentPanelInstance = &filamentPanel{
			CommonPanel: CreateCommonPanel("FilamentPanel", ui),
		}
		filamentPanelInstance.initialize()
		go filamentPanelInstance.consumeStateUpdates(ui.Printer.GetStateUpdates())
	}

	return filamentPanelInstance
}

func (this *filamentPanel) initialize() {
	defer this.Initialize()

	// Create the step buttons first, since they are needed by some of the other controls.
	this.flowRateStepButton = uiWidgets.CreateFlowRateStepButton(this.UI.Client, this.UI.Config, 1, nil)
	this.amountToExtrudeStepButton = uiWidgets.CreateAmountToExtrudeStepButton(this.UI.Config, 2, nil)

	// First row
	this.filamentExtrudeButton = uiWidgets.CreateFilamentExtrudeButton(
		this.UI.window,
		this.UI.Client,
		this.UI.Config,
		this.amountToExtrudeStepButton,
		this.flowRateStepButton,
		true,
	)
	this.Grid().Attach(this.filamentExtrudeButton, 0, 0, 1, 1)

	this.Grid().Attach(this.flowRateStepButton, 1, 0, 1, 1)

	this.Grid().Attach(this.amountToExtrudeStepButton, 2, 0, 1, 1)

	this.filamentRetractButton = uiWidgets.CreateFilamentExtrudeButton(
		this.UI.window,
		this.UI.Client,
		this.UI.Config,
		this.amountToExtrudeStepButton,
		this.flowRateStepButton,
		false,
	)
	this.Grid().Attach(this.filamentRetractButton, 3, 0, 1, 1)

	// Second row
	// this.filamentLoadButton = uiWidgets.CreateFilamentLoadButton(
	// 	this.UI.window,
	// 	this.UI.Client,
	// 	this.flowRateStepButton,
	// 	this.selectExtruderStepButton,
	// 	true,
	// 	int(this.UI.Settings.FilamentInLength),
	// )
	// this.Grid().Attach(this.filamentLoadButton, 0, 1, 1, 1)

	this.temperatureStatusBox = uiWidgets.CreateTemperatureStatusBox(this.UI.Client, this.UI.Config)
	this.Grid().Attach(this.temperatureStatusBox, 1, 1, 2, 1)

	// this.filamentUnloadButton = uiWidgets.CreateFilamentLoadButton(
	// 	this.UI.window,
	// 	this.UI.Client,
	// 	this.flowRateStepButton,
	// 	this.selectExtruderStepButton,
	// 	false,
	// 	int(this.UI.Settings.FilamentOutLength),
	// )
	// this.Grid().Attach(this.filamentUnloadButton, 3, 1, 1, 1)

	// Third row
	this.addTemperatureButton()
}

func (this *filamentPanel) consumeStateUpdates(ch chan domain.PrinterState) {
	for state := range ch {
		if state.IsConnectedToPrinter {
			glib.IdleAdd(func() {
				this.temperatureStatusBox.UpdateTemperatureData(state.Temperature)
			})
		}
	}
}

func (this *filamentPanel) showTemperaturePanel() {
	this.UI.GoToPanel(GetTemperaturePanelInstance(this.UI))
}

func (this *filamentPanel) addTemperatureButton() {
	this.temperatureButton = utils.MustButtonImageStyle(this.UI.Config, "Temperature", "heat-up.svg", "color1", this.showTemperaturePanel)
	this.Grid().Attach(this.temperatureButton, 0, 2, 1, 1)
}
