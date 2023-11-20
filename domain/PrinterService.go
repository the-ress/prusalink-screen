package domain

import (
	// "os"
	// "time"
	// "strconv"
	// "sync"

	// "github.com/gotk3/gotk3/glib"

	"sync"
	"time"

	"github.com/the-ress/prusalink-screen/logger"
	"github.com/the-ress/prusalink-screen/octoprintApis"
	"github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
	"github.com/the-ress/prusalink-screen/utils"
)

// const MAX_CONNECTION_ATTEMPTS = 8
// const MAX_CONNECTION_ATTEMPTS = 4

type subscriberRecord struct {
	ch chan *dataModels.FullStateResponse
}

type PrinterService struct {
	subscribersLock sync.RWMutex
	lock            sync.RWMutex

	client            *octoprintApis.Client
	backgroundTask    *utils.BackgroundTask
	fullStateResponse *dataModels.FullStateResponse

	subscribers []*subscriberRecord
}

func NewPrinterService(client *octoprintApis.Client) *PrinterService {
	instance := &PrinterService{
		client:      client,
		subscribers: []*subscriberRecord{},
	}

	instance.createBackgroundTask()

	return instance
}

func (this *PrinterService) createBackgroundTask() {
	logger.TraceEnter("PrinterService.createBackgroundTask()")

	// Default timeout of 10 seconds.
	duration := utils.GetExperimentalFrequency(10, "EXPERIMENTAL_OCTO_PRINT_RESPONSE_MANGER_UPDATE_FREQUENCY")
	this.backgroundTask = utils.CreateBackgroundTaskAnyThread(duration, this.updateState)
	this.backgroundTask.Start()

	logger.TraceLeave("PrinterService.createBackgroundTask()")
}

func (this *PrinterService) updateStateAfterTemperatureChange() {
	time.Sleep(time.Second * 2)
	this.updateState()
}

func (this *PrinterService) updateState() {
	logger.TraceEnter("PrinterService.updateState()")

	connectionManager := utils.GetConnectionManagerInstance(this.client)
	if connectionManager.IsConnected() != true {
		// If not connected, do nothing and leave.
		logger.TraceLeave("PrinterService.updateState()")
		return
	}

	fullStateResponse, err := (&octoprintApis.FullStateRequest{}).Do(this.client)
	if err != nil {
		logger.LogError("PrinterService.updateState()", "Do(FullStateRequest)", err)

		// connectionManager.ReInitializeConnectionState()
		// GoToConnectionPanel(this.UI)

		logger.TraceLeave("PrinterService.updateState()")
		return
	}

	if fullStateResponse == nil || fullStateResponse.Temperature.CurrentTemperatureData == nil {
		logger.Error("PrinterService.updateState() - fullStateResponse.Temperature.CurrentTemperatureData is invalid")

		// connectionManager.ReInitializeConnectionState()
		// GoToConnectionPanel(this.UI)

		logger.TraceLeave("PrinterService.updateState()")
		return
	}

	this.setFullStateRespone(fullStateResponse)
	this.notifySubscribers(fullStateResponse)

	logger.TraceLeave("PrinterService.update()")
}

func (this *PrinterService) setFullStateRespone(fullStateResponse *dataModels.FullStateResponse) {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.fullStateResponse = fullStateResponse
}

func (this *PrinterService) GetFullStateRespone() *dataModels.FullStateResponse {
	this.lock.Lock()
	defer this.lock.Unlock()

	return this.fullStateResponse
}

func (this *PrinterService) notifySubscribers(fullStateResponse *dataModels.FullStateResponse) {
	this.subscribersLock.Lock()
	defer this.subscribersLock.Unlock()

	for _, r := range this.subscribers {
		r.ch <- fullStateResponse
	}
}

func (this *PrinterService) GetStateUpdates() chan *dataModels.FullStateResponse {
	fullStateResponse := this.GetFullStateRespone()

	ch := make(chan *dataModels.FullStateResponse, 10)
	if fullStateResponse != nil {
		ch <- fullStateResponse
	}

	this.subscribersLock.Lock()
	defer this.subscribersLock.Unlock()

	this.subscribers = append(this.subscribers, &subscriberRecord{ch: ch})
	return ch
}

func (this *PrinterService) SetHotendTemperature(target float64) error {
	cmd := &octoprintApis.ToolTargetRequest{
		Targets: map[string]float64{"tool0": target},
	}
	err := cmd.Do(this.client)

	if err == nil {
		go this.updateStateAfterTemperatureChange()
	}

	return err
}

func (this *PrinterService) SetBedTemperature(target float64) error {
	cmd := &octoprintApis.BedTargetRequest{Target: target}
	err := cmd.Do(this.client)

	if err == nil {
		go this.updateStateAfterTemperatureChange()
	}

	return err
}

func (this *PrinterService) IsConnected() bool {
	// TODO: should this be named IsFullyConnected()?
	connectionManager := utils.GetConnectionManagerInstance(this.client)
	return connectionManager.IsConnected()
}

// func (this *connectionManager) ReInitializeConnectionState() {
// 	// this.IsRunning = true
// 	this.ConnectAttempts = 0
// 	this.IsConnectedToOctoPrint = false
// 	this.IsConnectedToPrinter = false
// }

// func (this *connectionManager) UpdateStatus() {
// 	logger.TraceEnter("ConnectionManager.UpdateStatus()")

// 	if this.IsConnected() != true {
// 		if this.ConnectAttempts > MAX_CONNECTION_ATTEMPTS {
// 			// verify
// 			this.ConnectAttempts++
// 		}

// 		logger.Debug("ConnectionManager.UpdateStatus() - about to call ConnectionRequest.Do()")
// 		t1 := time.Now()
// 		connectionResponse, err := (&octoprintApis.ConnectionRequest{}).Do(this.Client)
// 		t2 := time.Now()
// 		logger.Debug("ConnectionManager.UpdateStatus() - finished calling ConnectionRequest.Do()")
// 		logger.Debugf("time elapsed: %q", t2.Sub(t1))

// 		if err != nil {
// 			logger.LogError("ConnectionManager.UpdateStatus()", "ConnectionRequest.Do()", err)
// 								// newUIState, splashMessage = this.getUiStateAndMessageFromError(err, newUIState, splashMessage)
// 								// logger.Debugf("ConnectionManager.UpdateStatus() - newUIState is now: %s", newUIState)
// 			this.IsConnectedToOctoPrint = false
// 			logger.Debug("ConnectionManager.UpdateStatus() - Connection state: IsConnectedToOctoPrint is now false")
// 			logger.TraceLeave("ConnectionManager.UpdateStatus()")
// 			return
// 		}

// 		logger.Debug("ConnectionManager.UpdateStatus() - ConnectionRequest.Do() succeeded")

// 		this.IsConnectedToOctoPrint = true

// 		jsonResponse, err := StructToJson(connectionResponse)
// 		if err != nil {
// 			logger.LogError("ConnectionManager.UpdateStatus()", "StructToJson()", err)
// 			// If there's an error here, it's with the serialization of the object to JSON.
// 			// This is just for debugging, so don't return if there's an issue, and just
// 			// carry on (and hopefully connectionResponse isn't corrupted)
// 		} else {
// 			logger.Debugf("ConnectionManager.UpdateStatus() - connectionResponse is: %s", jsonResponse)
// 		}

// 		/*
// 		Example JSON response:
// 		{
// 			"Current": {
// 				"state": "Operational",
// 				"port": "/dev/ttyACM0",
// 				"baudrate": 115200,
// 				"printerProfile": "_default"
// 			},
// 			"Options": {
// 				"ports": [
// 					"/dev/ttyACM0"
// 				],
// 				"baudrates": [
// 					250000,
// 					230400,
// 					115200,
// 					57600,
// 					38400,
// 					19200,
// 					9600
// 				],
// 				"printerProfiles": [
// 					{
// 						"id": "_default",
// 						"name": "name-of-the-printer"
// 					}
// 				],
// 				"portPreference": "",
// 				"baudratePreference": 0,
// 				"printerProfilePreference": "_default",
// 				"autoconnect": false
// 			}
// 		}
// 		*/

// 		printerConnectionState := connectionResponse.Current.State
// 		if printerConnectionState.IsOffline() || printerConnectionState.IsError() {
// 			this.IsConnectedToPrinter = false
// 		} else {
// 			this.IsConnectedToPrinter = true
// 		}
// 	}

// 	logger.TraceLeave("ConnectionManager.UpdateStatus()")
// }

// func (this *connectionManager) IsConnected() bool {
// 	// TODO: should this be named IsFullyConnected?
// 	return this.IsConnectedToOctoPrint == true && this.IsConnectedToPrinter == true;
// }
