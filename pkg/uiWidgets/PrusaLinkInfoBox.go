package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type PrusaLinkInfoBox struct {
	*SystemInfoBox
}

func NewPrusaLinkInfoBox(
	config *config.ScreenConfig,
	logoWidth int,
) *PrusaLinkInfoBox {
	logoHeight := logoWidth

	base := NewSystemInfoBox(
		uiUtils.MustImageFromFileWithSize(config, "logos/prusa-link.svg", logoWidth, logoHeight),
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
