package uiWidgets

import (

	// "strconv"
	// "strings"

	"github.com/the-ress/prusalink-screen/prusaLinkApis"
	// "github.com/the-ress/prusalink-screen/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/utils"
)

func CreateSelectExtruderStepButton(
	client *prusaLinkApis.Client,
	config *utils.ScreenConfig,
	includeBed bool,
	colorVariation int,
	clicked func(),
) *SelectToolStepButton {
	var steps []Step

	step := Step{
		"Extruder",
		utils.GetExtruderFileName(),
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
