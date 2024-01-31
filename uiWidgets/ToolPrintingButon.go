package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/utils"
)

type ToolPrintingButton struct {
	*gtk.Button
}

func CreateToolPrintingButton(
	config *utils.ScreenConfig,
	index int,
) *ToolPrintingButton {
	imageFileName := ToolImageFileName(index)
	instance := &ToolPrintingButton{
		Button: utils.MustButtonImageUsingFilePath(config, "", imageFileName, nil),
	}

	ctx, _ := instance.GetStyleContext()
	ctx.AddClass("printing-state")

	return instance
}
