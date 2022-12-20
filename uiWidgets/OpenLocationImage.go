package uiWidgets

import (
	"fmt"

	"github.com/Z-Bolt/OctoScreen/utils"
	"github.com/gotk3/gotk3/gtk"
)

func CreateOpenLocationImage(
	index int,
	buttonWidth int,
	buttonHeight int,
	pixbufCache *utils.PixbufCache,
) *gtk.Image {
	colorClass := fmt.Sprintf("color%d", (index%4)+1)

	return CreateActionImage("open.svg", buttonWidth, buttonHeight, colorClass, pixbufCache)
}
