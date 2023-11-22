package ui

import (
	"fmt"
	// "math"
	// "os"
	// "strconv"
	"strings"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/the-ress/prusalink-screen/domain"
	"github.com/the-ress/prusalink-screen/logger"
	"github.com/the-ress/prusalink-screen/octoprintApis"
	"github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
	"github.com/the-ress/prusalink-screen/uiWidgets"
	"github.com/the-ress/prusalink-screen/utils"
)

type printStatusPanel struct {
	CommonPanel

	PrintWasCanceled bool

	// Tools
	tool0Button *uiWidgets.ToolPrintingButton
	tool1Button *uiWidgets.ToolPrintingButton
	tool2Button *uiWidgets.ToolPrintingButton
	tool3Button *uiWidgets.ToolPrintingButton
	tool4Button *uiWidgets.ToolPrintingButton
	bedButton   *uiWidgets.ToolPrintingButton

	// Statistics/Info
	fileLabelWithImage     *utils.LabelWithImage
	timeLabelWithImage     *utils.LabelWithImage
	timeLeftLabelWithImage *utils.LabelWithImage
	// layerLabelWithImage	*utils.LabelWithImage
	// The info for the current / total layers is not available
	// See https://community.octoprint.org/t/layer-number-and-total-layers-from-api/8005/4
	// and https://docs.octoprint.org/en/master/api/datamodel.html#sec-api-datamodel-jobs-job
	// Darn.

	// Progress
	progressBar *gtk.ProgressBar

	// Toolbar buttons
	pauseButton     *gtk.Button
	cancelButton    *gtk.Button
	controlButton   *gtk.Button
	completedButton *gtk.Button
}

var printStatusPanelInstance *printStatusPanel

func getPrintStatusPanelInstance(ui *UI) *printStatusPanel {
	if printStatusPanelInstance == nil {
		printStatusPanelInstance = &printStatusPanel{
			CommonPanel:      CreateTopLevelCommonPanel("PrintStatusPanel", ui),
			PrintWasCanceled: false,
		}
		printStatusPanelInstance.initialize()
		go printStatusPanelInstance.consumeStateUpdates(ui.Printer.GetStateUpdates())
	}

	return printStatusPanelInstance
}

func GoToPrintStatusPanel(ui *UI) {
	instance := getPrintStatusPanelInstance(ui)
	instance.progressBar.SetText("0%")
	ui.GoToPanel(instance)
}

func (this *printStatusPanel) initialize() {
	defer this.Initialize()

	this.createToolButtons()
	this.Grid().Attach(this.createInfoBox(), 2, 0, 2, 1)
	this.Grid().Attach(this.createProgressBar(), 2, 1, 2, 1)
	this.createToolBarButtons()
}

func (this *printStatusPanel) createToolButtons() {
	this.tool0Button = uiWidgets.CreateToolPrintingButton(0)
	this.bedButton = uiWidgets.CreateToolPrintingButton(-1)
	this.Grid().Attach(this.tool0Button, 0, 0, 2, 1)
	this.Grid().Attach(this.bedButton, 0, 1, 2, 1)
}

func (this *printStatusPanel) createInfoBox() *gtk.Box {
	this.fileLabelWithImage = utils.MustLabelWithImage("file-gcode.svg", "")
	ctx, _ := this.fileLabelWithImage.GetStyleContext()
	ctx.AddClass("printing-status-label")

	this.timeLabelWithImage = utils.MustLabelWithImage("time.svg", "Print time:")
	ctx, _ = this.timeLabelWithImage.GetStyleContext()
	ctx.AddClass("printing-status-label")

	this.timeLeftLabelWithImage = utils.MustLabelWithImage("time.svg", "Print time left:")
	ctx, _ = this.timeLeftLabelWithImage.GetStyleContext()
	ctx.AddClass("printing-status-label")

	// this.layerLabelWithImage = utils.MustLabelWithImage("time.svg", "")
	// ctx, _ = this.layerLabelWithImage.GetStyleContext()
	// ctx.AddClass("printing-status-label")

	infoBox := utils.MustBox(gtk.ORIENTATION_VERTICAL, 5)
	infoBox.SetHAlign(gtk.ALIGN_START)
	infoBox.SetHExpand(true)
	infoBox.SetVExpand(true)
	infoBox.SetVAlign(gtk.ALIGN_CENTER)
	infoBox.Add(this.fileLabelWithImage)
	infoBox.Add(this.timeLabelWithImage)
	infoBox.Add(this.timeLeftLabelWithImage)
	// infoBox.Add(this.layerLabelWithImage)

	return infoBox
}

func (this *printStatusPanel) createProgressBar() *gtk.ProgressBar {
	this.progressBar = utils.MustProgressBar()
	this.progressBar.SetShowText(true)
	this.progressBar.SetMarginTop(10)
	this.progressBar.SetMarginEnd(this.Scaled(20))
	this.progressBar.SetVAlign(gtk.ALIGN_CENTER)
	this.progressBar.SetVExpand(true)

	ctx, _ := this.progressBar.GetStyleContext()
	ctx.AddClass("printing-progress-bar")

	return this.progressBar
}

func (this *printStatusPanel) createToolBarButtons() {
	this.pauseButton = utils.MustButtonImageUsingFilePath(
		"Pause",
		"pause.svg",
		this.handlePauseClicked,
	)
	this.Grid().Attach(this.pauseButton, 1, 2, 1, 1)

	this.cancelButton = utils.MustButtonImageStyle(
		"Cancel",
		"stop.svg",
		"color-warning-sign-yellow",
		this.handleCancelClicked,
	)
	this.Grid().Attach(this.cancelButton, 2, 2, 1, 1)

	this.controlButton = utils.MustButtonImageStyle(
		"Control",
		"printing-control.svg",
		"color3",
		this.handleControlClicked,
	)
	this.Grid().Attach(this.controlButton, 3, 2, 1, 1)

	this.completedButton = utils.MustButtonImageStyle(
		"Completed",
		"complete.svg",
		"color3",
		this.handleCompleteClicked,
	)
	this.Grid().Attach(this.completedButton, 1, 2, 3, 1)
}

func (this *printStatusPanel) update(state domain.PrinterState) {
	this.updateToolTemperature(state.Temperature)

	if state.Job == nil {
		logger.Debug("PrintStatusPanel.update() - No job")
		return
	}

	logger.Debugf("PrintStatusPanel.update() - job.State is %s", state.Job.State)

	this.updateStates(state.Job)
	this.updateInfoBox(state.Job)
	this.updateProgress(state.Job)
	this.updateToolBarButtons(state.Job)
}

func (this *printStatusPanel) consumeStateUpdates(ch chan domain.PrinterState) {
	for state := range ch {
		if state.IsConnectedToPrinter {
			glib.IdleAdd(func() {
				this.update(state)
			})
		}
	}
}

func (this *printStatusPanel) updateStates(job *dataModels.JobResponse) {
	if job.State == "STOPPED" {
		this.PrintWasCanceled = true
	}

	if job.State == "PRINTING" || job.State == "FINISHED" {
		if float64(job.TimeRemaining) <= 0.0 {
			// Special case for handling the buttons
			this.pauseButton.SetSensitive(false)
			this.pauseButton.Hide()

			this.cancelButton.SetSensitive(false)
			this.cancelButton.Hide()

			this.controlButton.SetSensitive(false)
			this.controlButton.Hide()

			this.completedButton.Show()
			this.completedButton.SetSensitive(true)

			this.progressBar.Hide()

			this.timeLeftLabelWithImage.Label.SetLabel("Print time left: 00:00:00")
		}
	}
}

func (this *printStatusPanel) updateToolTemperature(temperature dataModels.TemperatureData) {
	this.tool0Button.SetLabel(utils.GetTemperatureDataString(temperature.Nozzle))
	this.bedButton.SetLabel(utils.GetTemperatureDataString(temperature.Bed))
}

func (this *printStatusPanel) updateInfoBox(job *dataModels.JobResponse) {
	logger.TraceEnter("PrintStatusPanel.updateInfoBox()")

	if job.State != "PRINTING" {
		logger.TraceLeave("PrintStatusPanel.updateInfoBox()")
		return
	}

	jobFileName := "<i>not-set</i>"
	if job.File.DisplayName != "" {
		jobFileName = job.File.DisplayName
		jobFileName = strings.Replace(jobFileName, ".gcode", "", -1)
		jobFileName = strings.Replace(jobFileName, ".gco", "", -1)
		jobFileName = utils.TruncateString(jobFileName, 20)
	}

	this.fileLabelWithImage.Label.SetLabel(jobFileName)

	var timeSpent string
	var timeLeft string
	if job.Progress == 100 {
		timeSpent = fmt.Sprintf("Completed in %s", time.Duration(job.TimePrinting)*time.Second)
		timeLeft = ""
	} else {
		timePrinting := time.Duration(job.TimePrinting) * time.Second
		timeRemaining := time.Duration(job.TimeRemaining) * time.Second
		timeSpent = fmt.Sprintf("Print time: %s", formattedDuration(timePrinting))
		if job.InaccurateEstimates {
			timeLeft = fmt.Sprintf("Print time left: ?%s", formattedDuration(timeRemaining))
		} else {
			timeLeft = fmt.Sprintf("Print time left: %s", formattedDuration(timeRemaining))
		}
	}

	if job.TimePrinting == 0 {
		timeSpent = "Warming up..."
	}

	this.timeLabelWithImage.Label.SetLabel(timeSpent)
	this.timeLeftLabelWithImage.Label.SetLabel(timeLeft)

	logger.TraceLeave("PrintStatusPanel.updateInfoBox()")
}

func (this *printStatusPanel) updateProgress(job *dataModels.JobResponse) {
	logger.TraceEnter("PrintStatusPanel.updateProgress()")
	logger.Debugf("PrintStatusPanel.updateProgress() - job.State is: '%s'", job.State)

	if job.State != "PRINTING" {
		this.progressBar.SetText(job.State)
		logger.TraceLeave("PrintStatusPanel.updateProgress()")
		return
	}

	progresBarFraction := job.Progress / 100.0
	this.progressBar.SetFraction(progresBarFraction)
	this.progressBar.SetText(fmt.Sprintf("%d%%", int64(job.Progress)))

	logger.TraceLeave("PrintStatusPanel.updateProgress()")
}

func (this *printStatusPanel) updateToolBarButtons(job *dataModels.JobResponse) {
	logger.TraceEnter("PrintStatusPanel.updateToolBarButtons()")
	logger.Debugf("PrintStatusPanel.updateToolBarButtons() - job.State is: '%s'", job.State)

	switch job.State {
	case "PRINTING":
		this.pauseButton.SetLabel("Pause")
		pauseImage := utils.MustImageFromFile("pause.svg")
		this.pauseButton.SetImage(pauseImage)
		this.pauseButton.SetSensitive(true)
		this.pauseButton.Show()

		this.cancelButton.SetSensitive(true)
		this.cancelButton.Show()

		this.controlButton.SetSensitive(true)
		this.controlButton.Show()

		this.completedButton.SetSensitive(false)
		this.completedButton.Hide()
		break

	case "PAUSED":
		this.pauseButton.SetLabel("Resume")
		pauseImage := utils.MustImageFromFile("resume.svg")
		this.pauseButton.SetImage(pauseImage)
		this.pauseButton.SetSensitive(true)
		this.pauseButton.Show()

		this.cancelButton.SetSensitive(true)
		this.cancelButton.Show()

		this.controlButton.SetSensitive(true)
		this.controlButton.Show()

		this.completedButton.SetSensitive(false)
		this.completedButton.Hide()
		break

	case "STOPPED":
		this.pauseButton.SetSensitive(false)
		this.pauseButton.Show()

		this.cancelButton.SetSensitive(false)
		this.cancelButton.Show()

		this.controlButton.SetSensitive(true)
		this.controlButton.Show()

		this.completedButton.SetSensitive(false)
		this.completedButton.Hide()
		break

	case "FINISHED":
		break

	case "ERROR":
		break

	default:
		logLevel := logger.LogLevel()
		if logLevel == "debug" {
			logger.Debugf("PrintStatusPanel.updateToolBarButtons() - unknown job.State: '%s'", job.State)
			logger.Panicf("PrintStatusPanel.updateToolBarButtons() - unknown job.State: '%s'", job.State)
		}
	}

	logger.TraceLeave("PrintStatusPanel.updateToolBarButtons()")
}

func (this *printStatusPanel) handlePauseClicked() {
	logger.TraceEnter("PrintStatusPanel.handlePauseClicked()")

	// TODO: is this needed?
	// defer this.updateTemperature()

	cmd := &octoprintApis.PauseRequest{Action: dataModels.Toggle}
	err := cmd.Do(this.UI.Client)
	if err != nil {
		logger.LogError("PrintStatusPanel.handlePauseClicked()", "Do(PauseRequest)", err)
		logger.TraceLeave("PrintStatusPanel.handlePauseClicked()")
		return
	}

	logger.TraceLeave("PrintStatusPanel.handlePauseClicked()")
}

func (this *printStatusPanel) handleCancelClicked() {
	userResponse := this.confirmCancelDialogBox(
		this.UI.window,
		"Are you sure you want to cancel the current print?",
	)

	if userResponse == gtk.RESPONSE_YES {
		this.cancelPrintJob()
	}
}

func (this *printStatusPanel) handleControlClicked() {
	this.UI.GoToPanel(GetPrintMenuPanelInstance(this.UI))
}

func (this *printStatusPanel) handleCompleteClicked() {
	this.UI.WaitingForUserToContinue = false
}

func (this *printStatusPanel) confirmCancelDialogBox(
	parentWindow *gtk.Window,
	message string,
) gtk.ResponseType {
	dialogBox := gtk.MessageDialogNewWithMarkup(
		parentWindow,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_YES_NO,
		"",
	)

	dialogBox.SetMarkup(utils.CleanHTML(message))
	defer dialogBox.Destroy()

	box, _ := dialogBox.GetContentArea()
	box.SetMarginStart(15)
	box.SetMarginEnd(15)
	box.SetMarginTop(15)
	box.SetMarginBottom(15)

	ctx, _ := dialogBox.GetStyleContext()
	ctx.AddClass("dialog")

	userResponse := dialogBox.Run()

	return userResponse
}

func (this *printStatusPanel) cancelPrintJob() {
	logger.TraceEnter("PrintStatusPanel.cancelPrintJob()")

	err := (&octoprintApis.CancelRequest{}).Do(this.UI.Client)
	if err == nil {
		this.pauseButton.SetSensitive(false)
		this.cancelButton.SetSensitive(false)
		this.controlButton.SetSensitive(false)
	} else {
		logger.LogError("PrintStatusPanel.cancelPrintJob()", "Do(CancelRequest)", err)
	}

	logger.TraceLeave("PrintStatusPanel.cancelPrintJob()")
}

func formattedDuration(duration time.Duration) string {
	hours := duration / time.Hour
	duration -= hours * time.Hour

	minutes := duration / time.Minute
	duration -= minutes * time.Minute

	seconds := duration / time.Second

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
