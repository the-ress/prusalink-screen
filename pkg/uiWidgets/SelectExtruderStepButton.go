package uiWidgets

import (

	// "strconv"
	// "strings"

	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
	// "github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
)

func CreateSelectExtruderStepButton(
	client *prusaLinkApis.Client,
	imageLoader *uiUtils.ImageLoader,
	includeBed bool,
	colorVariation int,
	clicked func(),
) *SelectToolStepButton {
	var steps []Step

	step := Step{
		"Extruder",
		uiUtils.ExtruderSvg,
		nil,
		"tool0",
	}

	steps = append(steps, step)

	if includeBed {
		steps = append(steps, Step{"Bed", uiUtils.BedSvg, nil, "bed"})
	}

	base := CreateStepButton(
		imageLoader,
		colorVariation,
		clicked,
		steps...,
	)

	instance := &SelectToolStepButton{
		StepButton: base,
	}

	return instance
}
