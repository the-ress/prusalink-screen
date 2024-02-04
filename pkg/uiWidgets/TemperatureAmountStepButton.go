package uiWidgets

import "github.com/the-ress/prusalink-screen/pkg/uiUtils"

type TemperatureAmountStepButton struct {
	*StepButton
}

func CreateTemperatureAmountStepButton(
	imageLoader *uiUtils.ImageLoader,
	colorVariation int,
	clicked func(),
) *TemperatureAmountStepButton {
	base := CreateStepButton(
		imageLoader,
		colorVariation,
		clicked,
		Step{"10°C", uiUtils.MoveStepSvg, nil, 10.0},
		Step{"20°C", uiUtils.MoveStepSvg, nil, 20.0},
		Step{"50°C", uiUtils.MoveStepSvg, nil, 50.0},
		Step{" 1°C", uiUtils.MoveStepSvg, nil, 1.0},
		Step{" 5°C", uiUtils.MoveStepSvg, nil, 5.0},
	)

	instance := &TemperatureAmountStepButton{
		StepButton: base,
	}

	return instance
}

func (this *TemperatureAmountStepButton) Value() float64 {
	return this.StepButton.Value().(float64)
}
