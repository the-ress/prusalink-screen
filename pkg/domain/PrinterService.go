package domain

import (
	// "os"
	// "time"
	// "strconv"
	// "sync"

	// "github.com/gotk3/gotk3/glib"

	"time"

	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

type PrinterService struct {
	client       *prusaLinkApis.Client
	stateManager *StateManager

	backgroundTask *utils.BackgroundTask
}

func NewPrinterService(client *prusaLinkApis.Client) *PrinterService {
	instance := &PrinterService{
		client:       client,
		stateManager: NewStateManager(client),
	}

	instance.createBackgroundTask()

	return instance
}

func (this *PrinterService) createBackgroundTask() {
	logger.TraceEnter("PrinterService.createBackgroundTask()")

	duration := utils.GetExperimentalFrequency(1, "EXPERIMENTAL_OCTO_PRINT_RESPONSE_MANGER_UPDATE_FREQUENCY")
	this.backgroundTask = utils.CreateBackgroundTaskAnyThread(duration, this.updateState)
	this.backgroundTask.Start()

	logger.TraceLeave("PrinterService.createBackgroundTask()")
}

func (this *PrinterService) updateStateAfterTemperatureChange() {
	time.Sleep(time.Second * 2)
	this.stateManager.Update()
}

func (this *PrinterService) updateState() {
	this.stateManager.Update()
}

func (this *PrinterService) GetState() PrinterState {
	return this.stateManager.GetState()
}

func (this *PrinterService) GetStateUpdates() chan PrinterState {
	return this.stateManager.GetUpdates()
}

func (this *PrinterService) SetHotendTemperature(target float64) error {
	cmd := &prusaLinkApis.ToolTargetRequest{
		Targets: map[string]float64{"tool0": target},
	}
	err := cmd.Do(this.client)

	if err == nil {
		go this.updateStateAfterTemperatureChange()
	}

	return err
}

func (this *PrinterService) SetBedTemperature(target float64) error {
	cmd := &prusaLinkApis.BedTargetRequest{Target: target}
	err := cmd.Do(this.client)

	if err == nil {
		go this.updateStateAfterTemperatureChange()
	}

	return err
}

func (this *PrinterService) IsConnected() bool {
	return this.stateManager.IsConnected()
}
