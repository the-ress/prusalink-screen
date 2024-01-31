package uiWidgets

import (
	// "fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

func CreatePrintImage(
	buttonWidth int,
	buttonHeight int,
	pixbufCache *utils.PixbufCache,
) *gtk.Image {
	return CreateActionImage("print.svg", buttonWidth, buttonHeight, "color-warning-sign-yellow", pixbufCache)
}
