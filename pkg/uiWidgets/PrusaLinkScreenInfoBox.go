package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

type PrusaLinkScreenInfoBox struct {
	*SystemInfoBox
}

func NewPrusaLinkScreenInfoBox(
	config *utils.ScreenConfig,
	version string,
) *PrusaLinkScreenInfoBox {
	base := NewSystemInfoBox(
		utils.MustImageFromFile(config, "logos/octoscreen-isometric-90%.png"),
		"PrusaLink Screen",
	)

	instance := &PrusaLinkScreenInfoBox{
		SystemInfoBox: base,
	}

	instance.SetText(version)

	return instance
}
