package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/utils"
)

type PrusaLinkScreenInfoBox struct {
	*SystemInfoBox
}

func NewPrusaLinkScreenInfoBox(
	version string,
) *PrusaLinkScreenInfoBox {
	base := NewSystemInfoBox(
		utils.MustImageFromFile("logos/octoscreen-isometric-90%.png"),
		"PrusaLink Screen",
	)

	instance := &PrusaLinkScreenInfoBox{
		SystemInfoBox: base,
	}

	instance.SetText(version)

	return instance
}
