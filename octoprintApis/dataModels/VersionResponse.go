package dataModels

// VersionResponse is the response from a job command.
type VersionResponse struct {
	API      string `json:"api"`
	Firmware string `json:"firmware"`
	Server   string `json:"server"`
}
