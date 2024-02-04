package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

func CreateActionImage(
	fileName uiUtils.ImageFileName,
	buttonWidth int,
	buttonHeight int,
	colorClass string,
	imageLoader *uiUtils.ImageLoader,
) *gtk.Image {
	image := imageLoader.MustGetImageWithSize(
		fileName,
		buttonWidth,
		buttonHeight,
	)

	imageStyleContext, _ := image.GetStyleContext()
	imageStyleContext.AddClass(colorClass)

	return image
}
