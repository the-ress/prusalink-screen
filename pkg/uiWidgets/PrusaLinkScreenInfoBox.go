package uiWidgets

import (
	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type PrusaLinkScreenInfoBox struct {
	*SystemInfoBox
}

func NewPrusaLinkScreenInfoBox(
	config *config.ScreenConfig,
	version string,
) *PrusaLinkScreenInfoBox {
	base := NewSystemInfoBox(
		uiUtils.MustImageFromFile(config, "logos/octoscreen-isometric-90%.png"),
		"PrusaLink Screen",
	)

	instance := &PrusaLinkScreenInfoBox{
		SystemInfoBox: base,
	}

	instance.SetText(version)

	return instance
}
