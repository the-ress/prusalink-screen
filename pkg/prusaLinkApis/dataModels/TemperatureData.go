package dataModels

type TemperatureData struct {
	Nozzle ToolTemperatureData `json:"tool0"`
	Bed    ToolTemperatureData `json:"bed"`
}
