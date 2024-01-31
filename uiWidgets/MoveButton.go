package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/logger"
	"github.com/the-ress/prusalink-screen/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/utils"
)

type MoveButton struct {
	*gtk.Button

	client                 *prusaLinkApis.Client
	amountToMoveStepButton *AmountToMoveStepButton
	axis                   dataModels.Axis
	direction              float64
}

func CreateMoveButton(
	client *prusaLinkApis.Client,
	config *utils.ScreenConfig,
	amountToMoveStepButton *AmountToMoveStepButton,
	label string,
	image string,
	axis dataModels.Axis,
	direction float64,
) *MoveButton {
	// A little bit of a "chicken or the egg" situation here.  Create the
	// instance first so there is a reference to the callback...
	instance := &MoveButton{
		Button:                 nil,
		client:                 client,
		amountToMoveStepButton: amountToMoveStepButton,
		axis:                   axis,
		direction:              direction,
	}
	base := MustPressedButton(config, label, image, instance.handlePressed, 200)
	// ... and then set the button
	instance.Button = base

	return instance
}

func (this *MoveButton) handlePressed() {
	distance := this.amountToMoveStepButton.Value() * this.direction
	cmd := &prusaLinkApis.PrintHeadJogRequest{}
	switch this.axis {
	case dataModels.XAxis:
		cmd.X = distance

	case dataModels.YAxis:
		cmd.Y = distance

	case dataModels.ZAxis:
		cmd.Z = distance
	}

	if err := cmd.Do(this.client); err != nil {
		logger.LogError("MoveButton.handlePressed()", "Do(PrintHeadJogRequest)", err)
		return
	}
}
