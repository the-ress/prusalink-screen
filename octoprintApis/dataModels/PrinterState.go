package dataModels

// PrinterState current state of the printer.
type PrinterState struct {
	Text  string       `json:"text"`
	Flags PrinterFlags `json:"flags"`
}

type PrinterFlags struct {
	Operations    bool `json:"operational"`
	Paused        bool `json:"paused"`
	Printing      bool `json:"printing"`
	SDReady       bool `json:"sdReady"`
	Error         bool `json:"error"`
	Ready         bool `json:"ready"`
	ClosedOnError bool `json:"closedOrError"`
}
