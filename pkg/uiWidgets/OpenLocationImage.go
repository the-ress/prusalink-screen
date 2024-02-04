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
	imageLoader *uiUtils.ImageLoader,
) *gtk.Image {
	colorClass := fmt.Sprintf("color%d", (index%4)+1)

	return CreateActionImage(uiUtils.OpenSvg, buttonWidth, buttonHeight, colorClass, imageLoader)
}
