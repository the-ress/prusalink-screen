package domain

import (
	"sync"

	"github.com/the-ress/prusalink-screen/logger"
	"github.com/the-ress/prusalink-screen/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/prusaLinkApis/dataModels"
)

type StateManager struct {
	client *prusaLinkApis.Client
	state  PrinterState

	subscribers     []*subscriberRecord
	subscribersLock sync.RWMutex
	lock            sync.RWMutex
}

type PrinterState struct {
	IsConnectedToPrusaLink bool
	PrusaLinkErrorMessage  string

	IsConnectedToPrinter bool
	PrinterErrorMessage  string

	PrusaConnectStatus struct {
		OK      bool
		Message string
	}

	Text        dataModels.PrinterStateText
	Temperature dataModels.TemperatureData
	Job         *dataModels.JobResponse
}

type subscriberRecord struct {
	ch chan PrinterState
}

func NewStateManager(client *prusaLinkApis.Client) *StateManager {
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
	// logger.Debug("StateManager.detectState() - about to call ConnectionRequest.Do()")
	// t1 := time.Now()
	// connectionResponse, err := (&prusaLinkApis.ConnectionRequest{}).Do(this.client)
	// t2 := time.Now()
	// logger.Debug("StateManager.detectState() - finished calling ConnectionRequest.Do()")
	// logger.Debugf("time elapsed: %q", t2.Sub(t1))

	statusResponse, err := (&prusaLinkApis.StatusRequest{}).Do(this.client)
	if err != nil {
		logger.LogError("StateManager.detectState()", "Do(StatusRequest)", err)
		return PrinterState{
			IsConnectedToPrusaLink: false,
			PrusaLinkErrorMessage:  err.Error(),
		}
	}

	if statusResponse == nil {
		logger.Error("StateManager.detectState() - statusResponse is nil")
		return PrinterState{
			IsConnectedToPrusaLink: false,
			PrusaLinkErrorMessage:  "Response is empty",
		}
	}

	if !statusResponse.Printer.StatusPrinter.OK {
		logger.Errorf("StateManager.detectState() - printer is not OK: %s", statusResponse.Printer.StatusPrinter.Message)
		logger.Debug("StateManager.detectState() - Connection state: IsConnectedToPrinter is now false")
		return PrinterState{
			IsConnectedToPrusaLink: true,
			IsConnectedToPrinter:   false,
			PrinterErrorMessage:    statusResponse.Printer.StatusPrinter.Message,
		}
	}

	temperature := dataModels.TemperatureData{
		Nozzle: dataModels.ToolTemperatureData{
			Actual: statusResponse.Printer.TempNozzle,
			Target: statusResponse.Printer.TargetNozzle,
		},
		Bed: dataModels.ToolTemperatureData{
			Actual: statusResponse.Printer.TempBed,
			Target: statusResponse.Printer.TargetBed,
		},
	}

	state := PrinterState{
		IsConnectedToPrusaLink: true,
		IsConnectedToPrinter:   true,

		Text:        statusResponse.Printer.State,
		Temperature: temperature,
		PrusaConnectStatus: struct {
			OK      bool
			Message string
		}(statusResponse.Printer.StatusConnect),
	}

	jobResponse, err := (&prusaLinkApis.JobRequest{}).Do(this.client)
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
