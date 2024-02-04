package ui

import (
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
	"github.com/the-ress/prusalink-screen/pkg/uiWidgets"
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
	this.amountToMoveStepButton = uiWidgets.CreateAmountToMoveStepButton(this.UI.ImageLoader, 1, nil)

	moveXMinusImage := this.UI.ImageLoader.MustGetImage(uiUtils.MoveXMinusSvg)
	moveXPlusImage := this.UI.ImageLoader.MustGetImage(uiUtils.MoveXPlusSvg)
	moveYPlusImage := this.UI.ImageLoader.MustGetImage(uiUtils.MoveYPlusSvg)
	moveYMinusImage := this.UI.ImageLoader.MustGetImage(uiUtils.MoveYMinusSvg)
	moveZPlusImage := this.UI.ImageLoader.MustGetImage(uiUtils.MoveZPlusSvg)
	moveZMinusImage := this.UI.ImageLoader.MustGetImage(uiUtils.MoveZMinusSvg)

	this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "X-", moveXMinusImage, dataModels.XAxis, -1), 0, 1, 1, 1)
	this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "X+", moveXPlusImage, dataModels.XAxis, 1), 2, 1, 1, 1)

	this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Y+", moveYPlusImage, dataModels.YAxis, 1), 1, 0, 1, 1)
	this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Y-", moveYMinusImage, dataModels.YAxis, -1), 1, 2, 1, 1)

	this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Z+", moveZPlusImage, dataModels.ZAxis, 1), 3, 0, 1, 1)
	this.Grid().Attach(uiWidgets.CreateMoveButton(this.UI.Client, this.amountToMoveStepButton, "Z-", moveZMinusImage, dataModels.ZAxis, -1), 3, 1, 1, 1)

	this.Grid().Attach(this.amountToMoveStepButton, 1, 1, 1, 1)
}
