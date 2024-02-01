package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type PrinterFirmwareInfoBox struct {
	*SystemInfoBox
}

func NewPrinterFirmwareInfoBox(
	config *config.ScreenConfig,
	logoWidth int,
) *PrinterFirmwareInfoBox {
	logoHeight := int(float64(logoWidth) * 1.25)

	base := NewSystemInfoBox(
		uiUtils.MustImageFromFileWithSize(config, "logos/logo-octoprint.png", logoWidth, logoHeight),
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
