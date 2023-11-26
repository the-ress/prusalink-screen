package ui

import (
	// "time"

	// "github.com/the-ress/prusalink-screen/interfaces"

	"os"
	"os/exec"
	"path/filepath"

	"github.com/the-ress/prusalink-screen/logger"
	"github.com/the-ress/prusalink-screen/octoprintApis"
	"github.com/the-ress/prusalink-screen/uiWidgets"
	"github.com/the-ress/prusalink-screen/utils"
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
	this.prusaLinkInfoBox = uiWidgets.NewPrusaLinkInfoBox(logoWidth)
	this.Grid().Attach(this.prusaLinkInfoBox, 0, 0, 1, 1)

	this.printerFirmwareInfoBox = uiWidgets.NewPrinterFirmwareInfoBox(logoWidth)
	this.Grid().Attach(this.printerFirmwareInfoBox, 1, 0, 1, 1)

	this.prusaConnectInfoBox = uiWidgets.NewPrusaConnectInfoBox()
	this.Grid().Attach(this.prusaConnectInfoBox, 2, 0, 1, 1)

	this.prusaLinkScreenInfoBox = uiWidgets.NewPrusaLinkScreenInfoBox(utils.OctoScreenVersion)
	this.Grid().Attach(this.prusaLinkScreenInfoBox, 3, 0, 1, 1)

	// Second row
	this.shutdownSystemButton = uiWidgets.NewSystemCommandButton(
		this.UI.window,
		"Shutdown System",
		"shutdown",
		"color-warning-sign-yellow",
		"You are about to shutdown the system.\n\nThis action may disrupt any ongoing print jobs (depending on your setup).",
		func() {
			logger.Info("Shutting down system")

			err := exec.Command("sudo", "systemctl", "poweroff").Run()
			if err != nil {
				logger.LogError("SystemCommandButton callback", "exec.Command", err)
			}
		},
	)
	this.Grid().Attach(this.shutdownSystemButton, 0, 1, 1, 1)

	this.rebootSystemButton = uiWidgets.NewSystemCommandButton(
		this.UI.window,
		"Reboot System",
		"reboot",
		"color-warning-sign-yellow",
		"You are about to reboot the system.\n\nThis action may disrupt any ongoing print jobs (depending on your setup).",
		func() {
			logger.Info("Rebooting system")

			err := exec.Command("sudo", "systemctl", "reboot").Run()
			if err != nil {
				logger.LogError("SystemCommandButton callback", "exec.Command", err)
			}
		},
	)
	this.Grid().Attach(this.rebootSystemButton, 1, 1, 1, 1)

	this.restartPrusaLinkButton = uiWidgets.NewSystemCommandButton(
		this.UI.window,
		"Restart PrusaLink",
		"restart",
		"color-warning-sign-yellow",
		"You are about to restart the PrusaLink server.\n\nThis action may disrupt any ongoing print jobs (depending on your setup).",
		func() {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				logger.LogError("SystemCommandButton callback", "os.UserHomeDir", err)
				return
			}

			logger.Info("Restarting PrusaLink")

			cmd := filepath.Join(homeDir, ".local/bin/prusalink")
			err = exec.Command(cmd, "-i", "restart").Run()
			if err != nil {
				logger.LogError("SystemCommandButton callback", "exec.Command", err)
			}
		},
	)
	this.Grid().Attach(this.restartPrusaLinkButton, 2, 1, 1, 1)

	this.restartPrusaLinkScreenButton = uiWidgets.NewSystemCommandButton(
		this.UI.window,
		"Restart PrusaLink Screen",
		"restart",
		"color-warning-sign-yellow",
		"You are about to restart the PrusaLink Screen UI.",
		func() {
			logger.Info("Restarting PrusaLink Screen")

			err := exec.Command("sudo", "systemctl", "restart", "prusalink-screen").Run()
			if err != nil {
				logger.LogError("SystemCommandButton callback", "Do(SystemExecuteCommandRequest)", err)
			}
		},
	)
	this.Grid().Attach(this.restartPrusaLinkScreenButton, 3, 1, 1, 1)

	// Third row
	this.systemInformationInfoBox = uiWidgets.CreateSystemInformationInfoBox(this.UI.window, this.UI.scaleFactor)
	this.Grid().Attach(this.systemInformationInfoBox, 0, 2, 3, 1)
}

func (this *systemPanel) refreshData() {
	prusaLinkVersion := "unknown"
	firmwareVersion := "unknown"
	prusaConnectStatus := "unknown"

	state := this.UI.Printer.GetState()
	if state.IsConnectedToPrusaLink {
		// Only call if we're connected to PrusaLink
		versionResponse, err := (&octoprintApis.VersionRequest{}).Do(this.UI.Client)
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
