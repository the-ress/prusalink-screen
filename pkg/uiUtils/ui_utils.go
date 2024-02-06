package uiUtils

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

// ****************************************************************************
// DialogBox Routines
func MustConfirmDialogBox(parent *gtk.Window, msg string, cb func()) {
	if yesNoDialogBox(parent, msg) == gtk.RESPONSE_YES {
		// glib.IdleAdd(func() { // Hide the confirmation window before continuing
		cb()
		// })
	}
}

func yesNoDialogBox(parentWindow *gtk.Window, message string) gtk.ResponseType {
	dialog := gtk.MessageDialogNewWithMarkup(
		parentWindow,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_YES_NO,
		"",
	)

	dialog.SetMarkup(utils.CleanHTML(message))
	defer dialog.Destroy()

	box, _ := dialog.GetContentArea()
	box.SetMarginStart(15)
	box.SetMarginEnd(15)
	box.SetMarginTop(15)
	box.SetMarginBottom(15)

	ctx, _ := dialog.GetStyleContext()
	ctx.AddClass("dialog")

	return dialog.Run()
}

func InfoMessageDialogBox(parentWindow *gtk.Window, message string) {
	messageDialogBox(parentWindow, gtk.MESSAGE_INFO, message)
}

func WarningMessageDialogBox(parentWindow *gtk.Window, message string) {
	messageDialogBox(parentWindow, gtk.MESSAGE_WARNING, message)
}

func ErrorMessageDialogBox(parentWindow *gtk.Window, message string) {
	messageDialogBox(parentWindow, gtk.MESSAGE_ERROR, message)
}

func messageDialogBox(parentWindow *gtk.Window, messageType gtk.MessageType, message string) {
	dialogBox := gtk.MessageDialogNewWithMarkup(
		parentWindow,
		gtk.DIALOG_MODAL,
		messageType,
		// gtk.BUTTONS_OK,
		gtk.BUTTONS_NONE,
		"",
	)

	dialogBox.AddButton("Continue", gtk.RESPONSE_OK)

	dialogBox.SetMarkup(utils.CleanHTML(message))
	defer dialogBox.Destroy()

	box, _ := dialogBox.GetContentArea()
	box.SetMarginStart(25)
	box.SetMarginEnd(25)
	box.SetMarginTop(20)
	box.SetMarginBottom(10)

	ctx, _ := dialogBox.GetStyleContext()
	ctx.AddClass("message")

	dialogBox.Run()
}

func RunWithWaitingBox(
	parentWindow *gtk.Window,
	message string,
	callback func(),
) {
	dialogBox := gtk.MessageDialogNewWithMarkup(
		parentWindow,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_NONE,
		"",
	)

	dialogBox.SetMarkup(utils.CleanHTML(message))
	defer dialogBox.Destroy()

	box, _ := dialogBox.GetContentArea()
	box.SetMarginStart(25)
	box.SetMarginEnd(25)
	box.SetMarginTop(20)
	box.SetMarginBottom(10)

	ctx, _ := dialogBox.GetStyleContext()
	ctx.AddClass("message")

	go func() {
		callback()
		glib.IdleAdd(func() {
			dialogBox.Close()
		})
	}()

	dialogBox.Run()
}

func EmptyTheContainer(container *gtk.Container) {
	children := container.GetChildren()
	defer children.Free()

	children.Foreach(func(i interface{}) {
		container.Remove(i.(gtk.IWidget))
	})
}
