package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/prusaLinkApis"
	// "github.com/the-ress/prusalink-screen/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/utils"
)

type FlowRateStepButton struct {
	*StepButton
	client *prusaLinkApis.Client
}

func CreateFlowRateStepButton(
	client *prusaLinkApis.Client,
	config *utils.ScreenConfig,
	colorVariation int,
	clicked func(),
) *FlowRateStepButton {
	base := CreateStepButton(
		config,
		colorVariation,
		clicked,
		Step{"Normal (100%)", "speed-normal.svg", nil, 100},
		Step{"Fast (125%)", "speed-fast.svg", nil, 125},
		Step{"Slow (75%)", "speed-slow.svg", nil, 75},
	)

	instance := &FlowRateStepButton{
		StepButton: base,
		client:     client,
	}

	return instance
}

func (this *FlowRateStepButton) Value() int {
	return this.StepButton.Value().(int)
}

func (this *FlowRateStepButton) SendChangeFlowRate() error {
	return utils.SetFlowRate(this.client, this.Value())
}
