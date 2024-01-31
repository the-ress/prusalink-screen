package uiWidgets

import (
	// "time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

type TemperatureStatusBox struct {
	*gtk.Box

	client      *prusaLinkApis.Client
	nozzleLabel *utils.LabelWithImage
	bedLabel    *utils.LabelWithImage
}

func CreateTemperatureStatusBox(
	client *prusaLinkApis.Client,
	config *utils.ScreenConfig,
) *TemperatureStatusBox {
	base := utils.MustBox(gtk.ORIENTATION_VERTICAL, 5)

	instance := &TemperatureStatusBox{
		Box:    base,
		client: client,
	}

	instance.SetVAlign(gtk.ALIGN_CENTER)
	instance.SetHAlign(gtk.ALIGN_CENTER)

	instance.nozzleLabel = utils.MustLabelWithImage(config, utils.GetNozzleFileName(), "")
	instance.Add(instance.nozzleLabel)

	instance.bedLabel = utils.MustLabelWithImage(config, "bed.svg", "")
	instance.Add(instance.bedLabel)

	return instance
}

func (this *TemperatureStatusBox) UpdateTemperatureData(temperatureData dataModels.TemperatureData) {
	this.nozzleLabel.Label.SetText(utils.GetTemperatureDataString(temperatureData.Nozzle))
	this.nozzleLabel.ShowAll()

	this.bedLabel.Label.SetText(utils.GetTemperatureDataString(temperatureData.Bed))
	this.bedLabel.ShowAll()
}
