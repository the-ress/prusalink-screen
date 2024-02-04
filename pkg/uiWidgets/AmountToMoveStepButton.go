package uiWidgets

import "github.com/the-ress/prusalink-screen/pkg/uiUtils"

type AmountToMoveStepButton struct {
	*StepButton
}

func CreateAmountToMoveStepButton(
	imageLoader *uiUtils.ImageLoader,
	colorVariation int,
	clicked func(),
) *AmountToMoveStepButton {
	base := CreateStepButton(
		imageLoader,
		colorVariation,
		clicked,
		Step{"10mm", uiUtils.MoveStepSvg, nil, 10.00},
		Step{"20mm", uiUtils.MoveStepSvg, nil, 20.00},
		Step{"50mm", uiUtils.MoveStepSvg, nil, 50.00},
		Step{"0.02mm", uiUtils.MoveStepSvg, nil, 0.02},
		Step{"0.1mm", uiUtils.MoveStepSvg, nil, 0.10},
		Step{" 1mm", uiUtils.MoveStepSvg, nil, 1.00},
	)

	instance := &AmountToMoveStepButton{
		StepButton: base,
	}

	return instance
}

func (this *AmountToMoveStepButton) Value() float64 {
	return this.StepButton.Value().(float64)
}
