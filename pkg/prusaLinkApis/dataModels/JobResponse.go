package dataModels

// JobResponse is the response from a job command.
type JobResponse struct {
	Id int `json:"id"`

	// PRINTING, PAUSED, FINISHED, STOPPED, ERROR
	State string `json:"state"`

	// Percents
	Progress float64 `json:"progress"`

	// Seconds
	TimeRemaining int `json:"time_remaining"`

	// Seconds
	TimePrinting int `json:"time_printing"`

	// Whether the time estimates are accurate or inaccurate
	InaccurateEstimates bool `json:"inaccurate_estimates"`

	// Job contains information regarding the target of the current print job.
	File JobFilePrint `json:"file"`
}

type JobFilePrint struct {
	// Short Filename
	Name string `json:"name"`

	// Long Filename
	DisplayName string `json:"display_name"`

	Path string `json:"path"`

	DisplayPath string `json:"display_path"`

	// Bytes
	Size int `json:"size"`

	// Timestamp in seconds
	Timestamp int `json:"m_timestamp"`

	Metadata PrintFileMetadata `json:"meta"`

	Refs PrintFileRefs `json:"refs"`
}

type PrintFileMetadata struct {
	FilamentType string `json:"filament_type"`

	// Degrees Celsius
	Temperature int `json:"temperature"`

	// Degrees Celsius
	BedTemperature int `json:"bed_temperature"`

	// Seconds
	EstimatedPrintTime int `json:"estimated_print_time"`

	// Milimeters
	MaxLayerZ float64 `json:"max_layer_z"`

	// Milimeters
	NozzleDiameter float64 `json:"nozzle_diameter"`

	// Percents, e.g. "15%"
	FillDensity string `json:"fill_density"`

	// Milimeters
	BrimWidth float64 `json:"brim_width"`

	// Milimeters
	LayerHeight float64 `json:"layer_height"`

	Ironing int `json:"ironing"`

	FilamentUsedInGrams float64 `json:"filament used [g]"`
}

type PrintFileRefs struct {
	Download  string `json:"download"`
	Icon      string `json:"icon"`
	Thumbnail string `json:"thumbnail"`
}
