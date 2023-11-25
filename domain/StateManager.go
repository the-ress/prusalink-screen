package domain

import (
	"sync"
	"time"

	"github.com/the-ress/prusalink-screen/logger"
	"github.com/the-ress/prusalink-screen/octoprintApis"
	"github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
)

type StateManager struct {
	client *octoprintApis.Client
	state  PrinterState

	subscribers     []*subscriberRecord
	subscribersLock sync.RWMutex
	lock            sync.RWMutex
}

type PrinterState struct {
	IsConnectedToPrusaLink bool
	IsConnectedToPrinter   bool
	Text                   dataModels.PrinterStateText
	Temperature            dataModels.TemperatureData
	Job                    *dataModels.JobResponse
}

type subscriberRecord struct {
	ch chan PrinterState
}

func NewStateManager(client *octoprintApis.Client) *StateManager {
	return &StateManager{
		client: client,
		state: PrinterState{
			IsConnectedToPrusaLink: false,
			IsConnectedToPrinter:   false,
		},
		subscribers: []*subscriberRecord{},
	}
}

func (this *StateManager) SetDisconnected() {
	this.publishState(
		PrinterState{
			IsConnectedToPrusaLink: false,
			IsConnectedToPrinter:   false,
		},
	)
}

func (this *StateManager) Update() {
	newState := this.detectState()
	this.publishState(newState)
}

func (this *StateManager) publishState(newState PrinterState) {
	this.setState(newState)
	this.notifySubscribers(newState)
}

func (this *StateManager) notifySubscribers(newState PrinterState) {
	this.subscribersLock.Lock()
	defer this.subscribersLock.Unlock()

	for _, r := range this.subscribers {
		r.ch <- newState
	}
}

func (this *StateManager) GetUpdates() chan PrinterState {
	ch := make(chan PrinterState, 10)
	ch <- this.GetState()

	this.subscribersLock.Lock()
	defer this.subscribersLock.Unlock()

	this.subscribers = append(this.subscribers, &subscriberRecord{ch: ch})
	return ch
}

func (this *StateManager) GetState() PrinterState {
	this.lock.Lock()
	defer this.lock.Unlock()

	return this.state
}

func (this *StateManager) setState(newState PrinterState) {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.state = newState
}

func (this *StateManager) detectState() PrinterState {
	// If OctoScreen is connected to PrusaLink,
	// and PrusaLink is connected to the printer,
	// don't bother checking again.
	// if !this.IsConnected() {
	// Continue on if OctoScreen isn't connected...

	logger.Debug("StateManager.detectState() - about to call ConnectionRequest.Do()")
	t1 := time.Now()
	connectionResponse, err := (&octoprintApis.ConnectionRequest{}).Do(this.client)
	t2 := time.Now()
	logger.Debug("StateManager.detectState() - finished calling ConnectionRequest.Do()")
	logger.Debugf("time elapsed: %q", t2.Sub(t1))

	if err != nil {
		logger.LogError("StateManager.detectState()", "ConnectionRequest.Do()", err)
		logger.Debug("StateManager.detectState() - Connection state: IsConnectedToPrusaLink is now false")
		return PrinterState{
			IsConnectedToPrusaLink: false,
			IsConnectedToPrinter:   false,
		}
	}

	printerConnectionState := connectionResponse.Current.State
	if printerConnectionState.IsError() {
		logger.Debug("StateManager.detectState() - Connection state: IsConnectedToPrinter is now false")
		return PrinterState{
			IsConnectedToPrusaLink: true,
			IsConnectedToPrinter:   false,
		}
	}

	fullStateResponse, err := (&octoprintApis.FullStateRequest{}).Do(this.client)
	if err != nil {
		logger.LogError("StateManager.detectState()", "Do(FullStateRequest)", err)
		return PrinterState{
			IsConnectedToPrusaLink: false,
			IsConnectedToPrinter:   false,
		}
	}

	if fullStateResponse == nil {
		logger.Error("StateManager.detectState() - fullStateResponse is nil")
		return PrinterState{
			IsConnectedToPrusaLink: false,
			IsConnectedToPrinter:   false,
		}
	}

	temperature := dataModels.TemperatureData{
		Nozzle: dataModels.ToolTemperatureData{
			Actual: fullStateResponse.Printer.TempNozzle,
			Target: fullStateResponse.Printer.TargetNozzle,
		},
		Bed: dataModels.ToolTemperatureData{
			Actual: fullStateResponse.Printer.TempBed,
			Target: fullStateResponse.Printer.TargetBed,
		},
	}

	state := PrinterState{
		IsConnectedToPrusaLink: true,
		IsConnectedToPrinter:   true,
		Text:                   fullStateResponse.Printer.State,
		Temperature:            temperature,
	}

	jobResponse, err := (&octoprintApis.JobRequest{}).Do(this.client)
	if err != nil {
		logger.LogError("StateManager.detectState()", "Do(JobRequest)", err)
	}
	state.Job = jobResponse

	return state
}

func (this *StateManager) IsConnected() bool {
	// TODO: should this be named IsFullyConnected()?
	return this.state.IsConnectedToPrusaLink && this.state.IsConnectedToPrinter
}
