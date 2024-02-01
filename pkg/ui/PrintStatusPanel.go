package ui

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"

	"github.com/the-ress/prusalink-screen/pkg/domain"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
	"github.com/the-ress/prusalink-screen/pkg/uiWidgets"
)

type printStatusPanel struct {
	CommonPanel

	// Tools
	nozzleButton *uiWidgets.ToolPrintingButton
	bedButton    *uiWidgets.ToolPrintingButton

	// Statistics/Info
	fileLabelWithImage     *uiUtils.LabelWithImage
	timeLabelWithImage     *uiUtils.LabelWithImage
	timeLeftLabelWithImage *uiUtils.LabelWithImage

	// Thumbnail
	thumbnailPath           string
	thumbnailBox            *gtk.Box
	thumbnailDisplayed      bool
	cancelThumbnailDownload context.CancelFunc

	// layerLabelWithImage	*uiUtils.LabelWithImage
	// The info for the current / total layers is not available
	// See https://community.octoprint.org/t/layer-number-and-total-layers-from-api/8005/4
	// and https://docs.octoprint.org/en/master/api/datamodel.html#sec-api-datamodel-jobs-job
	// Darn.

	// Progress
	progressBar *gtk.ProgressBar

	// Toolbar buttons
	pauseButton     *gtk.Button
	resumeButton    *gtk.Button
	cancelButton    *gtk.Button
	controlButton   *gtk.Button
	completedButton *gtk.Button

	currentJobId int
}

var printStatusPanelInstance *printStatusPanel

func getPrintStatusPanelInstance(ui *UI) *printStatusPanel {
	if printStatusPanelInstance == nil {
		printStatusPanelInstance = &printStatusPanel{
			CommonPanel:  CreateTopLevelCommonPanel("PrintStatusPanel", ui),
			currentJobId: -1,
		}
		printStatusPanelInstance.initialize()
		go printStatusPanelInstance.consumeStateUpdates(ui.Printer.GetStateUpdates())
	}

	return printStatusPanelInstance
}

func (this *printStatusPanel) initialize() {
	defer this.Initialize()

	this.createToolButtons()
	this.createThumbnailBox()
	this.Grid().Attach(this.createInfoBox(), 2, 0, 2, 1)
	this.Grid().Attach(this.createProgressBar(), 2, 1, 2, 1)
	this.createToolBarButtons()
}

func (this *printStatusPanel) createToolButtons() {
	this.nozzleButton = uiWidgets.CreateToolPrintingButton(this.UI.Config, 0)
	this.bedButton = uiWidgets.CreateToolPrintingButton(this.UI.Config, -1)
	this.Grid().Attach(this.nozzleButton, 0, 0, 2, 1)
	this.Grid().Attach(this.bedButton, 0, 1, 2, 1)
}

func (this *printStatusPanel) createThumbnailBox() {
	this.thumbnailBox = uiWidgets.CreateVerticalLayoutBox()

	this.thumbnailBox.SetMarginTop(this.Scaled(10))
	this.thumbnailBox.SetMarginBottom(this.Scaled(10))
	this.thumbnailBox.SetMarginStart(this.Scaled(10))
	this.thumbnailBox.SetMarginEnd(this.Scaled(10))

	this.thumbnailBox.SetVisible(false)
	this.Grid().Attach(this.thumbnailBox, 0, 0, 2, 2)
}

func (this *printStatusPanel) createInfoBox() *gtk.Box {
	this.fileLabelWithImage = uiUtils.MustLabelWithImage(this.UI.Config, "file-gcode.svg", "")
	ctx, _ := this.fileLabelWithImage.GetStyleContext()
	ctx.AddClass("printing-status-label")
	this.fileLabelWithImage.Label.SetEllipsize(pango.ELLIPSIZE_END)
	this.fileLabelWithImage.Label.SetMarginEnd(this.Scaled(20))

	this.timeLabelWithImage = uiUtils.MustLabelWithImage(this.UI.Config, "time.svg", "Print time:")
	ctx, _ = this.timeLabelWithImage.GetStyleContext()
	ctx.AddClass("printing-status-label")

	this.timeLeftLabelWithImage = uiUtils.MustLabelWithImage(this.UI.Config, "time.svg", "Print time left:")
	ctx, _ = this.timeLeftLabelWithImage.GetStyleContext()
	ctx.AddClass("printing-status-label")

	// this.layerLabelWithImage = uiUtils.MustLabelWithImage(this.UI.Config, "time.svg", "")
	// ctx, _ = this.layerLabelWithImage.GetStyleContext()
	// ctx.AddClass("printing-status-label")

	infoBox := uiUtils.MustBox(gtk.ORIENTATION_VERTICAL, 5)
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
	this.progressBar = uiUtils.MustProgressBar()
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
	this.pauseButton = uiUtils.MustButtonImageUsingFilePath(
		this.UI.Config,
		"Pause",
		"pause.svg",
		this.handlePauseClicked,
	)
	this.Grid().Attach(this.pauseButton, 1, 2, 1, 1)

	this.resumeButton = uiUtils.MustButtonImageUsingFilePath(
		this.UI.Config,
		"Resume",
		"resume.svg",
		this.handleResumeClicked,
	)
	this.Grid().Attach(this.resumeButton, 1, 2, 1, 1)

	this.cancelButton = uiUtils.MustButtonImageStyle(
		this.UI.Config,
		"Cancel",
		"stop.svg",
		"color-warning-sign-yellow",
		this.handleCancelClicked,
	)
	this.Grid().Attach(this.cancelButton, 2, 2, 1, 1)

	this.controlButton = uiUtils.MustButtonImageStyle(
		this.UI.Config,
		"Control",
		"printing-control.svg",
		"color3",
		this.handleControlClicked,
	)
	this.Grid().Attach(this.controlButton, 3, 2, 1, 1)

	this.completedButton = uiUtils.MustButtonImageStyle(
		this.UI.Config,
		"Done",
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

	this.currentJobId = state.Job.Id

	logger.Debugf("PrintStatusPanel.update() - state.Text is %s", state.Text)

	this.updateInfoBox(state)
	this.updateThumbnail(state)
	this.updateProgress(state)
	this.updateToolBarButtons(state.Text)

	this.thumbnailBox.SetVisible(this.thumbnailDisplayed)
	this.nozzleButton.SetVisible(!this.thumbnailDisplayed)
	this.bedButton.SetVisible(!this.thumbnailDisplayed)
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

func (this *printStatusPanel) updateToolTemperature(temperature dataModels.TemperatureData) {
	this.nozzleButton.SetLabel(uiUtils.GetTemperatureDataString(temperature.Nozzle))
	this.bedButton.SetLabel(uiUtils.GetTemperatureDataString(temperature.Bed))
}

func (this *printStatusPanel) updateInfoBox(state domain.PrinterState) {
	job := state.Job

	jobFileName := "<i>not-set</i>"
	if job.File.DisplayName != "" {
		jobFileName = job.File.DisplayName
		jobFileName = strings.Replace(jobFileName, ".gcode", "", -1)
		jobFileName = strings.Replace(jobFileName, ".gco", "", -1)
	}

	this.fileLabelWithImage.Label.SetLabel(jobFileName)

	var timeSpent string
	var timeLeft string

	if job.Progress == 100 && state.Text == dataModels.FINISHED {
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

	if state.Text == dataModels.PRINTING && job.Progress < 100 && job.TimePrinting == 0 {
		timeSpent = "Warming up..."
	}

	this.timeLabelWithImage.Label.SetLabel(timeSpent)
	this.timeLeftLabelWithImage.Label.SetLabel(timeLeft)
	if timeLeft != "" {
		this.timeLeftLabelWithImage.Show()
	} else {
		this.timeLeftLabelWithImage.Hide()
	}
}

func (this *printStatusPanel) updateThumbnail(state domain.PrinterState) {
	job := state.Job

	if this.thumbnailPath == job.File.Refs.Thumbnail {
		return // No change
	}

	if this.cancelThumbnailDownload != nil {
		this.cancelThumbnailDownload()
		this.cancelThumbnailDownload = nil
	}

	this.thumbnailPath = job.File.Refs.Thumbnail

	ctx, cancel := context.WithCancel(context.Background())
	this.cancelThumbnailDownload = cancel
	go this.downloadThumbnail(ctx, this.thumbnailPath)
}

func (this *printStatusPanel) downloadThumbnail(
	ctx context.Context,
	thumbnailPath string,
) {
	logger.Debugf("PrintStatusPanel.downloadThumbnail() - thumbnailPath is %s", thumbnailPath)

	if thumbnailPath == "" {
		glib.IdleAddPriority(glib.PRIORITY_LOW, func() {
			this.displayThumbnail(nil)
		})
	}

	imageBuffer, imageFromUrlErr := (&prusaLinkApis.ThumbnailRequest{Path: thumbnailPath}).Do(this.UI.Client)
	if imageFromUrlErr != nil {
		logger.Error("PrintStatusPanel.downloadThumbnail() - error from ThumbnailRequest:", imageFromUrlErr)
		glib.IdleAddPriority(glib.PRIORITY_LOW, func() {
			this.displayThumbnail(nil)
		})
		return
	}

	glib.IdleAddPriority(glib.PRIORITY_LOW, func() {
		select {
		case <-ctx.Done():
			return // Abort
		default:
		}

		this.displayThumbnail(imageBuffer)
	})

}

func (this *printStatusPanel) displayThumbnail(imageBuffer []byte) {
	uiUtils.EmptyTheContainer(&this.thumbnailBox.Container)

	if imageBuffer == nil {
		this.thumbnailBox.SetVisible(false)
		this.nozzleButton.SetVisible(true)
		this.bedButton.SetVisible(true)
		this.thumbnailDisplayed = false
		return
	}

	previewImage, err :=
		uiUtils.ImageFromBufferAtSize(imageBuffer, this.thumbnailBox.GetAllocatedWidth(), this.thumbnailBox.GetAllocatedHeight())

	if err != nil {
		logger.Error("FilesPreviewSubRow.createPreviewThumbnail() - error from ImageFromBuffer:", err)
		this.thumbnailBox.SetVisible(false)
		this.nozzleButton.SetVisible(true)
		this.bedButton.SetVisible(true)
		this.thumbnailDisplayed = false
		return
	}

	previewImage.SetHExpand(true)
	previewImage.SetVExpand(true)

	this.thumbnailBox.Add(previewImage)
	this.thumbnailBox.ShowAll()

	this.thumbnailBox.SetVisible(true)
	this.nozzleButton.SetVisible(false)
	this.bedButton.SetVisible(false)
	this.thumbnailDisplayed = true
}

func (this *printStatusPanel) updateProgress(state domain.PrinterState) {
	logger.Debugf("PrintStatusPanel.updateProgress() - state.Text is: '%s'", state.Text)

	if state.Text == dataModels.FINISHED {
		this.progressBar.Hide()
		return
	}
	this.progressBar.Show()

	if state.Text != dataModels.PRINTING {
		this.progressBar.SetText(state.Text.String())
		return
	}

	progresBarFraction := state.Job.Progress / 100.0
	this.progressBar.SetFraction(progresBarFraction)
	this.progressBar.SetText(fmt.Sprintf("%d%%", int64(state.Job.Progress)))
}

func (this *printStatusPanel) updateToolBarButtons(stateText dataModels.PrinterStateText) {
	logger.Debugf("PrintStatusPanel.updateToolBarButtons() - stateText is: '%s'", stateText)

	switch stateText {
	case dataModels.PRINTING:
		this.pauseButton.SetSensitive(true)
		this.pauseButton.Show()

		this.resumeButton.SetSensitive(false)
		this.resumeButton.Hide()

		this.cancelButton.SetSensitive(true)
		this.cancelButton.Show()

		this.controlButton.SetSensitive(true)
		this.controlButton.Show()

		this.completedButton.SetSensitive(false)
		this.completedButton.Hide()

	case dataModels.PAUSED:
		this.pauseButton.SetSensitive(false)
		this.pauseButton.Hide()

		this.resumeButton.SetSensitive(true)
		this.resumeButton.Show()

		this.cancelButton.SetSensitive(true)
		this.cancelButton.Show()

		this.controlButton.SetSensitive(true)
		this.controlButton.Show()

		this.completedButton.SetSensitive(false)
		this.completedButton.Hide()

	default:
		this.pauseButton.SetSensitive(false)
		this.pauseButton.Hide()

		this.resumeButton.SetSensitive(false)
		this.resumeButton.Hide()

		this.cancelButton.SetSensitive(false)
		this.cancelButton.Hide()

		this.controlButton.SetSensitive(false)
		this.controlButton.Hide()

		this.completedButton.Show()
		this.completedButton.SetSensitive(true)
	}
}

func (this *printStatusPanel) handlePauseClicked() {
	cmd := &prusaLinkApis.JobPauseRequest{JobId: this.currentJobId}
	err := cmd.Do(this.UI.Client)
	if err != nil {
		logger.LogError("PrintStatusPanel.handlePauseClicked()", "Do(PauseRequest)", err)
		return
	}
}

func (this *printStatusPanel) handleResumeClicked() {
	cmd := &prusaLinkApis.JobResumeRequest{JobId: this.currentJobId}
	err := cmd.Do(this.UI.Client)
	if err != nil {
		logger.LogError("PrintStatusPanel.handleResumeClicked()", "Do(ResumeRequest)", err)
		return
	}
}

func (this *printStatusPanel) handleCancelClicked() {
	uiUtils.MustConfirmDialogBox(
		this.UI.window,
		"Are you sure you want to cancel the current print?",
		func() {
			this.cancelPrintJob()
		},
	)
}

func (this *printStatusPanel) handleControlClicked() {
	this.UI.GoToPanel(GetPrintMenuPanelInstance(this.UI))
}

func (this *printStatusPanel) handleCompleteClicked() {
	this.UI.SetFinishedIdle()
}

func (this *printStatusPanel) cancelPrintJob() {
	err := (&prusaLinkApis.JobStopRequest{JobId: this.currentJobId}).Do(this.UI.Client)
	if err == nil {
		this.pauseButton.SetSensitive(false)
		this.resumeButton.SetSensitive(false)
		this.cancelButton.SetSensitive(false)
		this.controlButton.SetSensitive(false)
	} else {
		logger.LogError("PrintStatusPanel.cancelPrintJob()", "Do(CancelRequest)", err)
	}
}

func formattedDuration(duration time.Duration) string {
	hours := duration / time.Hour
	duration -= hours * time.Hour

	minutes := duration / time.Minute
	duration -= minutes * time.Minute

	return fmt.Sprintf("%02d:%02d", hours, minutes)

	// seconds := duration / time.Second
	// return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
