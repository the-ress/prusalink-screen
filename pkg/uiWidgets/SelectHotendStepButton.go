package uiWidgets

import (

	// "strconv"
	// "strings"

	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
	// "github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
)

func CreateSelectHotendStepButton(
	client *prusaLinkApis.Client,
	config *config.ScreenConfig,
	includeBed bool,
	colorVariation int,
	clicked func(),
) *SelectToolStepButton {
	var steps []Step

	step := Step{
		"Hotend",
		uiUtils.GetHotendFileName(),
		nil,
		"tool0",
	}

	steps = append(steps, step)

	if includeBed {
		steps = append(steps, Step{"Bed", "bed.svg", nil, "bed"})
	}

	base := CreateStepButton(
		config,
		colorVariation,
		clicked,
		steps...,
	)

	instance := &SelectToolStepButton{
		StepButton: base,
	}

	return instance
}
