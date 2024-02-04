package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type PrusaLinkScreenInfoBox struct {
	*SystemInfoBox
}

func NewPrusaLinkScreenInfoBox(
	imageLoader *uiUtils.ImageLoader,
	version string,
) *PrusaLinkScreenInfoBox {
	base := NewSystemInfoBox(
		imageLoader.MustGetImage(uiUtils.PrusaLinkScreenLogoPng),
		"PrusaLink Screen",
	)

	instance := &PrusaLinkScreenInfoBox{
		SystemInfoBox: base,
	}

	instance.SetText(version)

	return instance
}
