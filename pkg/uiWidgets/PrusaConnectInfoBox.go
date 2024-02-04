package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type PrusaConnectInfoBox struct {
	*SystemInfoBox
}

func NewPrusaConnectInfoBox(
	config *config.ScreenConfig,
	logoWidth int,
) *PrusaConnectInfoBox {
	logoHeight := logoWidth

	base := NewSystemInfoBox(
		uiUtils.MustImageFromFileWithSize(config, "logos/prusa-connect.png", logoWidth, logoHeight),
		"PrusaConnect",
	)

	instance := &PrusaConnectInfoBox{
		SystemInfoBox: base,
	}

	return instance
}

func (this *PrusaConnectInfoBox) SetStatus(status string) {
	this.SetText(status)
}
