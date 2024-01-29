package prusaLinkApis

import "fmt"

// "io"
// "github.com/the-ress/prusalink-screen/prusaLinkApis/dataModels"

// JobStopRequest stops the current print job.
type JobStopRequest struct {
	JobId int
}

// Do sends an API request and returns an error if any.
func (cmd *JobStopRequest) Do(c *Client) error {
	uri := fmt.Sprintf("%s/%d", JobApiUri, cmd.JobId)
	_, err := c.doJsonRequest("DELETE", uri, nil, JobToolErrors, true)
	return err
}
