package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type TemperatureStatusBox struct {
	*gtk.Box

	client      *prusaLinkApis.Client
	nozzleLabel *uiUtils.LabelWithImage
	bedLabel    *uiUtils.LabelWithImage
}

func CreateTemperatureStatusBox(
	client *prusaLinkApis.Client,
	imageLoader *uiUtils.ImageLoader,
) *TemperatureStatusBox {
	base := uiUtils.MustBox(gtk.ORIENTATION_VERTICAL, 5)

	instance := &TemperatureStatusBox{
		Box:    base,
		client: client,
	}

	instance.SetVAlign(gtk.ALIGN_CENTER)
	instance.SetHAlign(gtk.ALIGN_CENTER)

	instance.nozzleLabel = uiUtils.MustLabelWithImage(imageLoader, uiUtils.NozzleSvg, "")
	instance.Add(instance.nozzleLabel)

	instance.bedLabel = uiUtils.MustLabelWithImage(imageLoader, uiUtils.BedSvg, "")
	instance.Add(instance.bedLabel)

	return instance
}

func (this *TemperatureStatusBox) UpdateTemperatureData(temperatureData dataModels.TemperatureData) {
	this.nozzleLabel.Label.SetText(uiUtils.GetTemperatureDataString(temperatureData.Nozzle))
	this.nozzleLabel.ShowAll()

	this.bedLabel.Label.SetText(uiUtils.GetTemperatureDataString(temperatureData.Bed))
	this.bedLabel.ShowAll()
}
