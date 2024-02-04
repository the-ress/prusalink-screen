package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type PrusaConnectInfoBox struct {
	*SystemInfoBox
}

func NewPrusaConnectInfoBox(
	imageLoader *uiUtils.ImageLoader,
	logoWidth int,
) *PrusaConnectInfoBox {
	logoHeight := logoWidth

	base := NewSystemInfoBox(
		imageLoader.MustGetImageWithSize(uiUtils.PrusaConnectLogoPng, logoWidth, logoHeight),
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
