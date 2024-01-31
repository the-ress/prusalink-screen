package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

type PrusaLinkInfoBox struct {
	*SystemInfoBox
}

func NewPrusaLinkInfoBox(
	config *utils.ScreenConfig,
	logoWidth int,
) *PrusaLinkInfoBox {
	logoHeight := int(float64(logoWidth) * 1.25)

	base := NewSystemInfoBox(
		utils.MustImageFromFileWithSize(config, "logos/logo-octoprint.png", logoWidth, logoHeight),
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
