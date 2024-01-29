package prusaLinkApis

import (
	"bytes"
	"encoding/json"

	// "fmt"
	"io"
	// "strings"
	// "github.com/the-ress/prusalink-screen/prusaLinkApis/dataModels"
)

// ToolExtrudeRequest extrudes the given amount of filament from the currently
// selected tool.
type ToolExtrudeRequest struct {
	// Amount is the amount of filament to extrude in mm. May be negative to
	// retract.
	Amount int `json:"amount"`
}

// Do sends an API request and returns an error if any.
func (cmd *ToolExtrudeRequest) Do(c *Client) error {
	buffer := bytes.NewBuffer(nil)
	if err := cmd.encode(buffer); err != nil {
		return err
	}

	_, err := c.doJsonRequest("POST", PrinterToolApiUri, buffer, PrintToolErrors, true)
	return err
}

func (cmd *ToolExtrudeRequest) encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(struct {
		Command string `json:"command"`
		ToolExtrudeRequest
	}{
		Command:            "extrude",
		ToolExtrudeRequest: *cmd,
	})
}
