package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
	// "github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
)

type FlowRateStepButton struct {
	*StepButton
	client *prusaLinkApis.Client
}

func CreateFlowRateStepButton(
	client *prusaLinkApis.Client,
	config *config.ScreenConfig,
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
	return uiUtils.SetFlowRate(this.client, this.Value())
}
