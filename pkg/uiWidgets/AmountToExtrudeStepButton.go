package uiWidgets

import "github.com/the-ress/prusalink-screen/pkg/uiUtils"

type AmountToExtrudeStepButton struct {
	*StepButton
}

func CreateAmountToExtrudeStepButton(
	imageLoader *uiUtils.ImageLoader,
	colorVariation int,
	clicked func(),
) *AmountToExtrudeStepButton {
	base := CreateStepButton(
		imageLoader,
		colorVariation,
		clicked,
		Step{" 20mm", uiUtils.MoveStepSvg, nil, 20},
		Step{" 50mm", uiUtils.MoveStepSvg, nil, 50},
		Step{"100mm", uiUtils.MoveStepSvg, nil, 100},
		Step{"  1mm", uiUtils.MoveStepSvg, nil, 1},
		Step{"  5mm", uiUtils.MoveStepSvg, nil, 5},
		Step{" 10mm", uiUtils.MoveStepSvg, nil, 10},
	)

	instance := &AmountToExtrudeStepButton{
		StepButton: base,
	}

	return instance
}

func (this *AmountToExtrudeStepButton) Value() int {
	return this.StepButton.Value().(int)
}
