package uiWidgets

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type SystemCommandButton struct {
	*gtk.Button
}

func NewSystemCommandButton(
	parentWindow *gtk.Window,
	config *config.ScreenConfig,
	name string,
	action string,
	style string,
	confirmation string,
	callback func(),
) *SystemCommandButton {
	confirmationMessage := fmt.Sprintf("%s\n\nDo you wish to proceed?", confirmation)

	base := uiUtils.MustButtonImageStyle(
		config,
		name,
		action+".svg",
		style,
		func() { uiUtils.MustConfirmDialogBox(parentWindow, confirmationMessage, callback) },
	)
	ctx, _ := base.GetStyleContext()
	ctx.AddClass("font-size-19")

	instance := &SystemCommandButton{
		Button: base,
	}

	return instance
}
