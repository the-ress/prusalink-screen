package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/utils"
)

type PrusaLinkInfoBox struct {
	*SystemInfoBox
}

func NewPrusaLinkInfoBox(
	logoWidth int,
) *PrusaLinkInfoBox {
	logoHeight := int(float64(logoWidth) * 1.25)

	base := NewSystemInfoBox(
		utils.MustImageFromFileWithSize("logos/logo-octoprint.png", logoWidth, logoHeight),
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
