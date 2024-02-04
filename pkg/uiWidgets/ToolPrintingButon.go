package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type ToolPrintingButton struct {
	*gtk.Button
}

func CreateToolPrintingButton(
	index int,
	imageLoader *uiUtils.ImageLoader,
) *ToolPrintingButton {
	imageFileName := ToolImageFileName(index)
	image := imageLoader.MustGetImage(imageFileName)

	instance := &ToolPrintingButton{
		Button: uiUtils.MustButtonImage("", image, nil),
	}

	ctx, _ := instance.GetStyleContext()
	ctx.AddClass("printing-state")

	return instance
}
