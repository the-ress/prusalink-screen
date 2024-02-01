package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

func CreatePrintImage(
	buttonWidth int,
	buttonHeight int,
	pixbufCache *uiUtils.PixbufCache,
) *gtk.Image {
	return CreateActionImage("print.svg", buttonWidth, buttonHeight, "color-warning-sign-yellow", pixbufCache)
}
