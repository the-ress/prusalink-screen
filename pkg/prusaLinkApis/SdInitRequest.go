package prusaLinkApis

// "bytes"
// "encoding/json"
// "fmt"
// "io"
// "strings"

// "github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"

// SdInitRequest initializes the printer’s SD card, making it available for use.
// This also includes an initial retrieval of the list of files currently stored
// on the SD card.
type SdInitRequest struct{}

// Do sends an API request and returns an error if any.
func (cmd *SdInitRequest) Do(c *Client) error {
	return doCommandRequest(c, PrinterSdApiUri, "init", PrintSdErrors)
}
