package dataModels

// FullStateResponse contains informantion about the current state of the printer.
type FullStateResponse struct {
	Printer StatusPrinter `json:"printer"`
}

// Telemetry info about printer, all values except state are optional
type StatusPrinter struct {
	// IDLE, BUSY, PRINTING, PAUSED, FINISHED, STOPPED, ERROR, ATTTENTION, READY
	State PrinterStateText `json:"state"`

	TempNozzle float64 `json:"temp_nozzle"`

	TargetNozzle float64 `json:"target_nozzle"`

	TempBed float64 `json:"temp_bed"`

	TargetBed float64 `json:"target_bed"`

	// Available only when printer is not moving
	AxisX float64 `json:"axis_x"`

	// Available only when printer is not moving
	AxisY float64 `json:"axis_y"`

	AxisZ float64 `json:"axis_z"`

	Flow int `json:"flow"`

	Speed int `json:"speed"`

	FanHotend int `json:"fan_hotend"`

	FanPrint int `json:"fan_print"`

	StatusPrinter struct {
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
	} `json:"status_printer"`

	StatusConnect struct {
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
	} `json:"status_connect"`
}
