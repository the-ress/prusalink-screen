package uiWidgets

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

func CreateOpenLocationImage(
	index int,
	buttonWidth int,
	buttonHeight int,
	pixbufCache *uiUtils.PixbufCache,
) *gtk.Image {
	colorClass := fmt.Sprintf("color%d", (index%4)+1)

	return CreateActionImage("open.svg", buttonWidth, buttonHeight, colorClass, pixbufCache)
}
