package uiWidgets

import (

	// "strconv"
	// "strings"

	"github.com/Z-Bolt/OctoScreen/octoprintApis"
	// "github.com/Z-Bolt/OctoScreen/octoprintApis/dataModels"
	"github.com/Z-Bolt/OctoScreen/utils"
)

func CreateSelectHotendStepButton(
	client *octoprintApis.Client,
	includeBed bool,
	colorVariation int,
	clicked func(),
) *SelectToolStepButton {
	var steps []Step

	step := Step{
		"Hotend",
		utils.GetHotendFileName(),
		nil,
		"tool0",
	}

	steps = append(steps, step)

	if includeBed {
		steps = append(steps, Step{"Bed", "bed.svg", nil, "bed"})
	}

	base := CreateStepButton(
		colorVariation,
		clicked,
		steps...,
	)

	instance := &SelectToolStepButton{
		StepButton: base,
	}

	return instance
}
