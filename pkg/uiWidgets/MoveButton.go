package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
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
	amountToMoveStepButton *AmountToMoveStepButton,
	label string,
	image *gtk.Image,
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
	base := MustPressedButton(label, image, instance.handlePressed, 200)
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
