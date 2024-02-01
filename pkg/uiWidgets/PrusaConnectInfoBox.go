package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type PrusaConnectInfoBox struct {
	*SystemInfoBox
}

func NewPrusaConnectInfoBox(config *config.ScreenConfig) *PrusaConnectInfoBox {
	base := NewSystemInfoBox(
		uiUtils.MustImageFromFile(config, "logos/puzzle-piece.png"),
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
