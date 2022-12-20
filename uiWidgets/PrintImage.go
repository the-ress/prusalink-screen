package uiWidgets

import (
	// "fmt"

	"github.com/Z-Bolt/OctoScreen/utils"
	"github.com/gotk3/gotk3/gtk"
)

func CreatePrintImage(
	buttonWidth int,
	buttonHeight int,
	pixbufCache *utils.PixbufCache,
) *gtk.Image {
	return CreateActionImage("print.svg", buttonWidth, buttonHeight, "color-warning-sign-yellow", pixbufCache)
}
