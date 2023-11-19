package ui

import (
	// "github.com/gotk3/gotk3/gtk"

	// "github.com/the-ress/prusalink-screen/interfaces"
	// "github.com/the-ress/prusalink-screen/octoprintApis"
	"github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
	"github.com/the-ress/prusalink-screen/uiWidgets"
	// "github.com/the-ress/prusalink-screen/utils"
)

type movePanel struct {
	CommonPanel
	amountToMoveStepButton *uiWidgets.AmountToMoveStepButton
}

var movePanelInstance *movePanel

func GetMovePanelInstance(
	ui *UI,
) *movePanel {
	if movePanelInstance == nil {
		instance := &movePanel{
			CommonPanel: CreateCommonPanel("MovePanel", ui),
		}
		instance.initialize()
		movePanelInstance = instance
	}

	return movePanelInstance
}

func (this *movePanel) initialize() {
	defer this.Initialize()

	// Create the step button first, since it is needed by some of the other controls.
	this.amountToMoveStepButton = uiWidgets.CreateAmountToMoveStepButton(1, nil)

	this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "X-", "move-x-.svg", dataModels.XAxis, -1), 0, 1, 1, 1)
	this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "X+", "move-x+.svg", dataModels.XAxis, 1), 2, 1, 1, 1)

	this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Y+", "move-y+.svg", dataModels.YAxis, 1), 1, 0, 1, 1)
	this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Y-", "move-y-.svg", dataModels.YAxis, -1), 1, 2, 1, 1)

	this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Z+", "move-z+.svg", dataModels.ZAxis, 1), 3, 0, 1, 1)
	this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Z-", "move-z-.svg", dataModels.ZAxis, -1), 3, 1, 1, 1)

	this.Grid().Attach(this.amountToMoveStepButton, 1, 1, 1, 1)
}
