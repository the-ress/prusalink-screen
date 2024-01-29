package dataModels

// ToolTemperatureData is temperature stats for a tool.
type ToolTemperatureData struct {
	// Actual current temperature.
	Actual float64 `json:"actual"`

	// Target temperature, may be nil if no target temperature is set.
	Target float64 `json:"target"`
}
