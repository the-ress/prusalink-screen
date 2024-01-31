package prusaLinkApis

// "bytes"
// "encoding/json"
// "fmt"
// "io"
// "strings"

// "github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"

// SdRefreshRequest Refreshes the list of files stored on the printer’s SD card.
type SdRefreshRequest struct{}

// Do sends an API request and returns an error if any.
func (cmd *SdRefreshRequest) Do(c *Client) error {
	return doCommandRequest(c, PrinterSdApiUri, "refresh", PrintSdErrors)
}
