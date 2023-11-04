package uiWidgets

import (
	"fmt"
	"sync"

	"github.com/Z-Bolt/OctoScreen/logger"
	"github.com/Z-Bolt/OctoScreen/octoprintApis"
	"github.com/Z-Bolt/OctoScreen/octoprintApis/dataModels"
	"github.com/Z-Bolt/OctoScreen/utils"
	"github.com/gotk3/gotk3/gtk"
)

func ToolImageFileName(
	index int,
) string {
	if index < 0 {
		return "bed.svg"
	} else if index == 0 {
		return "hotend.svg"
	} else {
		return fmt.Sprintf("hotend-%d.svg", index)
	}
}

func ToolName(
	index int,
) string {
	if index < 0 {
		return "bed"
	} else if index == 0 {
		return "tool0"
	} else {
		return fmt.Sprintf("tool%d", index-1)
	}
}

type ToolButton struct {
	*gtk.Button
	sync.RWMutex

	isHeating bool
	tool      string
	printer   *octoprintApis.Client
}

func CreateToolButton(
	index int,
	printer *octoprintApis.Client,
) *ToolButton {
	imageFileName := ToolImageFileName(index)
	toolName := ToolName(index)

	instance := &ToolButton{
		Button:  utils.MustButtonImageUsingFilePath("", imageFileName, nil),
		tool:    toolName,
		printer: printer,
	}

	instance.Connect("clicked", instance.clicked)

	return instance
}

func (this *ToolButton) UpdateStatus(heating bool) {
	ctx, _ := this.GetStyleContext()
	if heating {
		ctx.AddClass("active")
	} else {
		ctx.RemoveClass("active")
	}

	this.isHeating = heating
}

func (this *ToolButton) SetTemperatures(temperatureData dataModels.TemperatureData) {
	text := utils.GetTemperatureDataString(temperatureData)
	this.SetLabel(text)
	this.UpdateStatus(temperatureData.Target > 0)
}

func (this *ToolButton) GetProfileTemperature() float64 {
	if this.tool == "bed" {
		return 75
	} else {
		return 220
	}
}

func (this *ToolButton) clicked() {
	defer func() {
		this.UpdateStatus(!this.isHeating)
	}()

	var (
		target float64
		err    error
	)

	if this.isHeating {
		target = 0.0
	} else {
		target = this.GetProfileTemperature()
	}

	if this.tool == "bed" {
		cmd := &octoprintApis.BedTargetRequest{Target: target}
		err = cmd.Do(this.printer)
		if err != nil {
			logger.LogError("ToolButton.clicked()", "Do(BedTargetRequest)", err)
		}
	} else {
		cmd := &octoprintApis.ToolTargetRequest{Targets: map[string]float64{this.tool: target}}
		err = cmd.Do(this.printer)
		if err != nil {
			logger.LogError("ToolButton.clicked()", "Do(ToolTargetRequest)", err)
		}
	}
}
