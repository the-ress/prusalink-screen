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
	uri := URIPrinter

	// log.Printf("TODO-Remove: StateRequest (FullStateResponse) uri is: %s", uri)
	//StateRequest uri is: %s /api/printer?history=true&limit=1&exclude=sd,state
	/*
		{
			"temperature": {
				"bed": {
					"actual": 26.9,
					"offset": 0,
					"target": 0.0
				},
				"history": [
					{
						"bed": {
							"actual": 26.9,
							"target": 0.0
						},
						"time": 1598235178,
						"tool0": {
							"actual": 35.4,
							"target": 0.0
						}
					}
				],
				"tool0": {
					"actual": 35.4,
					"offset": 0,
					"target": 0.0
				}
			}
		}
	*/

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
