package dataModels

import "encoding/json"

type PrinterStateText struct {
	placeholder string
}

func (this PrinterStateText) String() string {
	return this.placeholder
}

func (this *PrinterStateText) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &this.placeholder)
}

var (
	IDLE      = PrinterStateText{"IDLE"}
	BUSY      = PrinterStateText{"BUSY"}
	PRINTING  = PrinterStateText{"PRINTING"}
	PAUSED    = PrinterStateText{"PAUSED"}
	FINISHED  = PrinterStateText{"FINISHED"}
	STOPPED   = PrinterStateText{"STOPPED"}
	ERROR     = PrinterStateText{"ERROR"}
	ATTENTION = PrinterStateText{"ATTENTION"}
	READY     = PrinterStateText{"READY"}

	IDLE_STATES     = []PrinterStateText{IDLE, BUSY, READY}
	PRINTING_STATES = []PrinterStateText{PRINTING, PAUSED, FINISHED, STOPPED}
	FINISHED_STATES = []PrinterStateText{FINISHED, STOPPED}
	ERROR_STATES    = []PrinterStateText{ERROR, ATTENTION}
)
