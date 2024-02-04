package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

func CreatePrintImage(
	buttonWidth int,
	buttonHeight int,
	imageLoader *uiUtils.ImageLoader,
) *gtk.Image {
	return CreateActionImage(uiUtils.PrintSvg, buttonWidth, buttonHeight, "color-warning-sign-yellow", imageLoader)
}
