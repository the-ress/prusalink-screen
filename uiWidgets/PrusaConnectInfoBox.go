package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/utils"
)

type PrusaConnectInfoBox struct {
	*SystemInfoBox
}

func NewPrusaConnectInfoBox() *PrusaConnectInfoBox {
	base := NewSystemInfoBox(
		utils.MustImageFromFile("logos/puzzle-piece.png"),
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
