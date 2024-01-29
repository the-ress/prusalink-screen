package prusaLinkApis

import (
	// "bytes"
	"encoding/json"

	"github.com/the-ress/prusalink-screen/prusaLinkApis/dataModels"
)

const URIPrinter = "/api/printer"

// StatusRequest retrieves the current state of the printer.
type StatusRequest struct{}

// Do sends an API request and returns the API response.
func (cmd *StatusRequest) Do(c *Client) (*dataModels.StatusResponse, error) {
	uri := "/api/v1/status"

	bytes, err := c.doJsonRequest("GET", uri, nil, PrintErrors, true)
	if err != nil {
		return nil, err
	}

	response := &dataModels.StatusResponse{}
	if err := json.Unmarshal(bytes, response); err != nil {
		return nil, err
	}

	return response, err
}
