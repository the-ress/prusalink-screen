package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type ToolPrintingButton struct {
	*gtk.Button
}

func CreateToolPrintingButton(
	config *config.ScreenConfig,
	index int,
) *ToolPrintingButton {
	imageFileName := ToolImageFileName(index)
	instance := &ToolPrintingButton{
		Button: uiUtils.MustButtonImageUsingFilePath(config, "", imageFileName, nil),
	}

	ctx, _ := instance.GetStyleContext()
	ctx.AddClass("printing-state")

	return instance
}
