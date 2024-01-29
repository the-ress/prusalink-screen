package prusaLinkApis

// "bytes"
// "encoding/json"
// "fmt"
// "io"
// "strings"

// "github.com/the-ress/prusalink-screen/prusaLinkApis/dataModels"

// SdReleaseRequest releases the SD card from the printer. The reverse operation
// to init. After issuing this command, the SD card won’t be available anymore,
// hence and operations targeting files stored on it will fail.
type SdReleaseRequest struct{}

// Do sends an API request and returns an error if any.
func (cmd *SdReleaseRequest) Do(c *Client) error {
	return doCommandRequest(c, PrinterSdApiUri, "release", PrintSdErrors)
}
