package uiWidgets

import (
	"fmt"
	"sync"

	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/domain"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

func ToolImageFileName(
	index int,
) uiUtils.ImageFileName {
	if index < 0 {
		return uiUtils.BedSvg
	} else {
		return uiUtils.HotendSvg
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
	printer   *domain.PrinterService
}

func CreateToolButton(
	index int,
	printer *domain.PrinterService,
	imageLoader *uiUtils.ImageLoader,
) *ToolButton {
	imageFileName := ToolImageFileName(index)
	toolName := ToolName(index)

	image := imageLoader.MustGetImage(imageFileName)

	instance := &ToolButton{
		Button:  uiUtils.MustButtonImage("", image, nil),
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

func (this *ToolButton) SetTemperatures(temperatureData dataModels.ToolTemperatureData) {
	text := uiUtils.GetTemperatureDataString(temperatureData)
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
	)

	if this.isHeating {
		target = 0.0
	} else {
		target = this.GetProfileTemperature()
	}

	if this.tool == "bed" {
		err := this.printer.SetBedTemperature(target)
		if err != nil {
			logger.LogError("ToolButton.clicked()", "Do(BedTargetRequest)", err)
		}
	} else {
		err := this.printer.SetHotendTemperature(target)
		if err != nil {
			logger.LogError("ToolButton.clicked()", "Do(ToolTargetRequest)", err)
		}
	}
}
