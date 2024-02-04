package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type PrinterFirmwareInfoBox struct {
	*SystemInfoBox
}

func NewPrinterFirmwareInfoBox(
	imageLoader *uiUtils.ImageLoader,
	logoWidth int,
) *PrinterFirmwareInfoBox {
	logoHeight := logoWidth

	base := NewSystemInfoBox(
		imageLoader.MustGetImageWithSize(uiUtils.PrusaPrinterLogoPng, logoWidth, logoHeight),
		"Firmware",
	)

	instance := &PrinterFirmwareInfoBox{
		SystemInfoBox: base,
	}

	return instance
}

func (this *PrinterFirmwareInfoBox) SetVersion(version string) {
	this.SetText(version)
}
