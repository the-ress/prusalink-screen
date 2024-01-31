package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"

	"github.com/the-ress/prusalink-screen/pkg/utils"
)

func CreateActionImage(
	imageFileName string,
	buttonWidth int,
	buttonHeight int,
	colorClass string,
	pixbufCache *utils.PixbufCache,
) *gtk.Image {
	image := pixbufCache.MustImageFromFileWithSize(
		imageFileName,
		buttonWidth,
		buttonHeight,
	)

	imageStyleContext, _ := image.GetStyleContext()
	imageStyleContext.AddClass(colorClass)

	return image
}
