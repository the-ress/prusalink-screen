package uiWidgets

import "github.com/the-ress/prusalink-screen/pkg/config"

type TemperatureAmountStepButton struct {
	*StepButton
}

func CreateTemperatureAmountStepButton(
	config *config.ScreenConfig,
	colorVariation int,
	clicked func(),
) *TemperatureAmountStepButton {
	base := CreateStepButton(
		config,
		colorVariation,
		clicked,
		Step{"10°C", "move-step.svg", nil, 10.0},
		Step{"20°C", "move-step.svg", nil, 20.0},
		Step{"50°C", "move-step.svg", nil, 50.0},
		Step{" 1°C", "move-step.svg", nil, 1.0},
		Step{" 5°C", "move-step.svg", nil, 5.0},
	)

	instance := &TemperatureAmountStepButton{
		StepButton: base,
	}

	return instance
}

func (this *TemperatureAmountStepButton) Value() float64 {
	return this.StepButton.Value().(float64)
}
