package ui

import (

	// "os"
	// "strconv"
	// "strings"
	// "sync"
	// "time"

	"fmt"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/the-ress/prusalink-screen/pkg/domain"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type connectionPanel struct {
	CommonPanel

	IsCheckingConnection bool
	printer              *domain.PrinterService

	// First row
	Logo *gtk.Image

	// Second row
	Label *gtk.Label

	// Third row
	ActionBar   *gtk.Box
	RetryButton *gtk.Button
}

var connectionPanelInstance *connectionPanel

func getConnectionPanelInstance(ui *UI, printer *domain.PrinterService) *connectionPanel {
	logger.TraceEnter("ConnectionPanel.getConnectionPanelInstance()")

	if connectionPanelInstance == nil {
		connectionPanelInstance = &connectionPanel{
			CommonPanel:          CreateCommonPanel("ConnectionPanel", ui),
			IsCheckingConnection: true,
			printer:              printer,
		}
		connectionPanelInstance.initialize()

		go connectionPanelInstance.consumeStateUpdates(ui.Printer.GetStateUpdates())
	}

	logger.TraceLeave("ConnectionPanel.getConnectionPanelInstance()")
	return connectionPanelInstance
}

func (this *connectionPanel) initialize() {
	logger.TraceEnter("ConnectionPanel.initialize()")

	_, windowHeight := this.UI.window.GetSize()
	unscaledLogo := uiUtils.MustImageFromFile(this.UI.Config, "logos/octoscreen-logo.svg")
	pixbuf := unscaledLogo.GetPixbuf()
	width := pixbuf.GetWidth()
	height := pixbuf.GetHeight()

	originalLogoWidth := 154.75
	originalLogoHeight := 103.75
	displayHeight := windowHeight / 2.0

	scaleFactor := float64(displayHeight) / originalLogoHeight
	displayWidth := int(originalLogoWidth * scaleFactor)
	displayHeight = int(originalLogoHeight * scaleFactor)

	this.Logo = uiUtils.MustImageFromFileWithSize(this.UI.Config, "logos/octoscreen-logo.svg", displayWidth, displayHeight)

	pixbuf.ScaleSimple(
		this.UI.scaleFactor*width,
		this.UI.scaleFactor*height,
		gdk.INTERP_NEAREST,
	)

	this.Label = uiUtils.MustLabel("Welcome to PrusaLink Screen")
	this.Label.SetHExpand(true)
	this.Label.SetHAlign(gtk.ALIGN_CENTER)
	this.Label.SetLineWrap(true)
	this.Label.SetMaxWidthChars(80)

	main := uiUtils.MustBox(gtk.ORIENTATION_VERTICAL, 15)
	main.SetHExpand(true)
	main.SetHAlign(gtk.ALIGN_CENTER)
	main.SetVExpand(true)
	main.SetVAlign(gtk.ALIGN_CENTER)

	main.Add(this.Logo)
	main.Add(this.Label)

	this.createActionBar()

	box := uiUtils.MustBox(gtk.ORIENTATION_VERTICAL, 0)
	box.Add(main)
	box.Add(this.ActionBar)
	this.Grid().Add(box)

	logger.TraceLeave("ConnectionPanel.initialize()")
}

func (this *connectionPanel) createActionBar() {
	logger.TraceEnter("ConnectionPanel.createActionBar()")

	this.ActionBar = uiUtils.MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	this.ActionBar.SetHAlign(gtk.ALIGN_END)

	refreshImage, err := this.UI.ImageLoader.GetImage(uiUtils.RefreshSvg)
	if err != nil {
		panic(err)
	}

	this.RetryButton = uiUtils.MustButtonImageStyle(refreshImage, "Retry", "color-none", this.initializeConnectionState)
	this.RetryButton.SetProperty("width-request", this.Scaled(100))

	this.ActionBar.Add(this.RetryButton)

	this.displayButtons(false)

	logger.TraceLeave("ConnectionPanel.createActionBar()")
}

func (this *connectionPanel) displayButtons(display bool) {
	retryButtonStyleContext, _ := this.RetryButton.GetStyleContext()
	if display {
		retryButtonStyleContext.RemoveClass("hidden")
		this.RetryButton.SetSensitive(true)
	} else {
		retryButtonStyleContext.AddClass("hidden")
		this.RetryButton.SetSensitive(false)
	}
}

func (this *connectionPanel) consumeStateUpdates(ch chan domain.PrinterState) {
	for state := range ch {
		glib.IdleAdd(func() {
			this.update(state)
		})
	}
}

func (this *connectionPanel) update(state domain.PrinterState) {
	var msg string
	if !state.IsConnectedToPrusaLink {
		msg = "Attempting to connect to PrusaLink"
		if state.PrusaLinkErrorMessage != "" {
			msg = fmt.Sprintf("%s\n%s", msg, state.PrusaLinkErrorMessage)
		}
		// if connectionManager.ConnectAttempts >= utils.MAX_CONNECTION_ATTEMPTS {
		// 	msg = fmt.Sprintf("Unable to connect to PrusaLink")
		// 	this.displayButtons(true)
		// } else if connectionManager.ConnectAttempts == 0 {
		// 	msg = fmt.Sprintf("Attempting to connect to PrusaLink")
		// } else {
		// 	msg = fmt.Sprintf("Attempting to connect to PrusaLink...%d", connectionManager.ConnectAttempts+1)
		// }
	} else if !state.IsConnectedToPrinter {
		msg = "Attempting to connect to the printer"
		if state.PrinterErrorMessage != "" {
			msg = fmt.Sprintf("%s\n%s", msg, state.PrinterErrorMessage)
		}
		// if connectionManager.ConnectAttempts >= utils.MAX_CONNECTION_ATTEMPTS {
		// 	msg = fmt.Sprintf("Unable to connect to the printer")
		// 	this.displayButtons(true)
		// } else if connectionManager.ConnectAttempts == 0 {
		// 	msg = fmt.Sprintf("Attempting to connect to the printer")
		// } else {
		// 	msg = fmt.Sprintf("Attempting to connect to the printer...%d", connectionManager.ConnectAttempts+1)
		// }
	} else {
		msg = "Connected"
	}

	this.Label.SetText(msg)
}

func (this *connectionPanel) initializeConnectionState() {
	this.displayButtons(false)

	this.Label.SetText("Attempting to connect to OctoPrint")
}

func (this *connectionPanel) showSystem() {
	this.UI.GoToPanel(GetSystemPanelInstance(this.UI))
}
