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
	logoHeight := int(float64(logoWidth) * 1.25)

	base := NewSystemInfoBox(
		uiUtils.MustImageFromFileWithSize(config, "logos/logo-octoprint.png", logoWidth, logoHeight),
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
