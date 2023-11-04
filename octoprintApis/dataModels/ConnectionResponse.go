package dataModels

// ConnectionResponse is the response from a connection command.
type ConnectionResponse struct {
	Current struct {
		// State current state of the connection.
		State ConnectionState `json:"state"`

		// Port to connect to.
		Port string `json:"port"`

		// BaudRate speed of the connection.
		BaudRate int `json:"baudrate"`
	}

	Options struct {
		// Ports list of available ports.
		Ports []string `json:"ports"`

		// BaudRates list of available speeds.
		BaudRates []int `json:"baudrates"`

		// PortPreference default port.
		PortPreference string `json:"portPreference"`

		// BaudRatePreference default speed.
		BaudRatePreference int `json:"baudratePreference"`

		// Autoconnect whether to automatically connect to the printer on
		// OctoPrint’s startup in the future.
		Autoconnect bool `json:"autoconnect"`
	}
}
