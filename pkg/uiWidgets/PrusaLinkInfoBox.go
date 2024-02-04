package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type PrusaLinkInfoBox struct {
	*SystemInfoBox
}

func NewPrusaLinkInfoBox(
	imageLoader *uiUtils.ImageLoader,
	logoWidth int,
) *PrusaLinkInfoBox {
	logoHeight := logoWidth

	base := NewSystemInfoBox(
		imageLoader.MustGetImageWithSize(uiUtils.PrusaLinkLogoSvg, logoWidth, logoHeight),
		"PrusaLink",
	)

	instance := &PrusaLinkInfoBox{
		SystemInfoBox: base,
	}

	return instance
}

func (this *PrusaLinkInfoBox) SetVersion(version string) {
	this.SetText(version)
}
