package ui

import (
	// "time"

	// "github.com/the-ress/prusalink-screen/pkg/interfaces"

	"fmt"
	"os/exec"

	"github.com/the-ress/prusalink-screen/pkg/common"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
	"github.com/the-ress/prusalink-screen/pkg/uiWidgets"
)

type systemPanel struct {
	CommonPanel

	// First row
	prusaLinkInfoBox       *uiWidgets.PrusaLinkInfoBox
	printerFirmwareInfoBox *uiWidgets.PrinterFirmwareInfoBox
	prusaConnectInfoBox    *uiWidgets.PrusaConnectInfoBox
	prusaLinkScreenInfoBox *uiWidgets.PrusaLinkScreenInfoBox

	// Second row
	systemInformationInfoBox *uiWidgets.SystemInformationInfoBox

	// Third row
	shutdownSystemButton         *uiWidgets.SystemCommandButton
	rebootSystemButton           *uiWidgets.SystemCommandButton
	restartPrusaLinkButton       *uiWidgets.SystemCommandButton
	restartPrusaLinkScreenButton *uiWidgets.SystemCommandButton
}

var systemPanelInstance *systemPanel = nil

func GetSystemPanelInstance(
	ui *UI,
) *systemPanel {
	if systemPanelInstance == nil {
		instance := &systemPanel{
			CommonPanel: CreateCommonPanel("SystemPanel", ui),
		}
		instance.initialize()
		instance.preShowCallback = instance.refreshData
		systemPanelInstance = instance
	}

	return systemPanelInstance
}

func (this *systemPanel) initialize() {
	defer this.Initialize()

	// First row
	logoWidth := this.Scaled(52)
	this.prusaLinkInfoBox = uiWidgets.NewPrusaLinkInfoBox(this.UI.ImageLoader, logoWidth)
	this.Grid().Attach(this.prusaLinkInfoBox, 0, 0, 1, 1)

	this.printerFirmwareInfoBox = uiWidgets.NewPrinterFirmwareInfoBox(this.UI.ImageLoader, logoWidth)
	this.Grid().Attach(this.printerFirmwareInfoBox, 1, 0, 1, 1)

	this.prusaConnectInfoBox = uiWidgets.NewPrusaConnectInfoBox(this.UI.ImageLoader, logoWidth)
	this.Grid().Attach(this.prusaConnectInfoBox, 2, 0, 1, 1)

	this.prusaLinkScreenInfoBox = uiWidgets.NewPrusaLinkScreenInfoBox(this.UI.ImageLoader, common.AppVersion)
	this.Grid().Attach(this.prusaLinkScreenInfoBox, 3, 0, 1, 1)

	// Second row
	this.shutdownSystemButton = uiWidgets.NewSystemCommandButton(
		this.UI.window,
		this.UI.ImageLoader,
		"Shutdown System",
		uiUtils.ShutdownSvg,
		"color-warning-sign-red",
		"You are about to shutdown the system.\n\nThis action may disrupt any ongoing print jobs (depending on your setup).",
		func() {
			logger.Info("Shutting down system")

			var err error = nil
			uiUtils.RunWithWaitingBox(this.UI.window, "Shutting down...", func() {
				err = exec.Command("systemctl", "poweroff").Run()
			})

			if err != nil {
				logger.LogError("SystemCommandButton callback", "exec.Command", err)
				errorMessage := fmt.Sprintf("Command failed:\n\n%s", err.Error())
				uiUtils.ErrorMessageDialogBox(this.UI.window, errorMessage)
				return
			}
		},
	)
	this.Grid().Attach(this.shutdownSystemButton, 0, 2, 1, 1)

	this.rebootSystemButton = uiWidgets.NewSystemCommandButton(
		this.UI.window,
		this.UI.ImageLoader,
		"Reboot System",
		uiUtils.RebootSvg,
		"color-warning-sign-red",
		"You are about to reboot the system.\n\nThis action may disrupt any ongoing print jobs (depending on your setup).",
		func() {
			logger.Info("Rebooting system")

			var err error = nil
			uiUtils.RunWithWaitingBox(this.UI.window, "Rebooting...", func() {
				err = exec.Command("systemctl", "reboot").Run()
			})

			if err != nil {
				logger.LogError("SystemCommandButton callback", "exec.Command", err)
				errorMessage := fmt.Sprintf("Command failed:\n\n%s", err.Error())
				uiUtils.ErrorMessageDialogBox(this.UI.window, errorMessage)
				return
			}
		},
	)
	this.Grid().Attach(this.rebootSystemButton, 1, 2, 1, 1)

	this.restartPrusaLinkButton = uiWidgets.NewSystemCommandButton(
		this.UI.window,
		this.UI.ImageLoader,
		"Restart PrusaLink",
		uiUtils.RestartSvg,
		"color-warning-sign-yellow-dark",
		"You are about to restart the PrusaLink server.\n\nThis action may disrupt any ongoing print jobs (depending on your setup).",
		func() {
			logger.Info("Restarting PrusaLink")

			command := fmt.Sprintf("%s -i restart", this.UI.Config.PrusaLinkExecutablePath)

			var err error = nil
			uiUtils.RunWithWaitingBox(this.UI.window, "Restarting PrusaLink...", func() {
				err = exec.Command("su", this.UI.Config.PrusaLinkUser, "-c", command).Run()
			})

			if err != nil {
				logger.LogError("SystemCommandButton callback", "exec.Command", err)
				errorMessage := fmt.Sprintf("Command failed:\n\n%s", err.Error())
				uiUtils.ErrorMessageDialogBox(this.UI.window, errorMessage)
				return
			}
		},
	)
	this.Grid().Attach(this.restartPrusaLinkButton, 0, 1, 1, 1)

	this.restartPrusaLinkScreenButton = uiWidgets.NewSystemCommandButton(
		this.UI.window,
		this.UI.ImageLoader,
		"Restart Screen",
		uiUtils.RestartSvg,
		"color-warning-sign-yellow-dark",
		"You are about to restart the PrusaLink Screen UI.",
		func() {
			logger.Info("Restarting PrusaLink Screen")

			var err error = nil
			uiUtils.RunWithWaitingBox(this.UI.window, "Restarting PrusaLink Screen...", func() {
				err = exec.Command("systemctl", "restart", "prusalink-screen").Run()
			})

			if err != nil {
				errorMessage := fmt.Sprintf("Command failed:\n\n%s", err.Error())
				uiUtils.ErrorMessageDialogBox(this.UI.window, errorMessage)
				logger.LogError("SystemCommandButton callback", "Do(SystemExecuteCommandRequest)", err)
				return
			}
		},
	)
	this.Grid().Attach(this.restartPrusaLinkScreenButton, 3, 1, 1, 1)

	// Third row
	this.systemInformationInfoBox = uiWidgets.CreateSystemInformationInfoBox(this.UI.window, this.UI.Config, this.UI.scaleFactor)
	this.Grid().Attach(this.systemInformationInfoBox, 1, 1, 2, 1)
}

func (this *systemPanel) refreshData() {
	prusaLinkVersion := "unknown"
	firmwareVersion := "unknown"
	prusaConnectStatus := "unknown"

	state := this.UI.Printer.GetState()
	if state.IsConnectedToPrusaLink {
		// Only call if we're connected to PrusaLink
		versionResponse, err := (&prusaLinkApis.VersionRequest{}).Do(this.UI.Client)
		if err != nil {
			logger.LogError("SystemPanel", "VersionRequest.Do()", err)
		} else if versionResponse != nil {
			prusaLinkVersion = versionResponse.Server
			firmwareVersion = versionResponse.Firmware
		}

		if state.PrusaConnectStatus.OK {
			if state.PrusaConnectStatus.Message == "OK" {
				prusaConnectStatus = "Linked"
			} else {
				prusaConnectStatus = "Not Linked"
			}
		} else {
			prusaConnectStatus = state.PrusaConnectStatus.Message
		}
	}

	this.prusaLinkInfoBox.SetVersion(prusaLinkVersion)
	this.printerFirmwareInfoBox.SetVersion(firmwareVersion)
	this.prusaConnectInfoBox.SetStatus(prusaConnectStatus)

	this.systemInformationInfoBox.Refresh()
}
