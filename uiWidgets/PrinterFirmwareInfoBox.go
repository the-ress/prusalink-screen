package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/utils"
)

type PrinterFirmwareInfoBox struct {
	*SystemInfoBox
}

func NewPrinterFirmwareInfoBox(
	logoWidth int,
) *PrinterFirmwareInfoBox {
	logoHeight := int(float64(logoWidth) * 1.25)

	base := NewSystemInfoBox(
		utils.MustImageFromFileWithSize("logos/logo-octoprint.png", logoWidth, logoHeight),
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
