package octoprintApis

import (
	// "bytes"
	"encoding/json"

	"github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
)

const URIPrinter = "/api/printer"

// FullStateRequest retrieves the current state of the printer.
type FullStateRequest struct{}

// Do sends an API request and returns the API response.
func (cmd *FullStateRequest) Do(c *Client) (*dataModels.FullStateResponse, error) {
	uri := "/api/v1/status"

	bytes, err := c.doJsonRequest("GET", uri, nil, PrintErrors, true)
	if err != nil {
		return nil, err
	}

	response := &dataModels.FullStateResponse{}
	if err := json.Unmarshal(bytes, response); err != nil {
		return nil, err
	}

	return response, err
}
