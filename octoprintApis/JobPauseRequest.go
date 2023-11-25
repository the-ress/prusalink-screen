package octoprintApis

import (
	"fmt"
)

// JobPauseRequest pauses/resumes/toggles the current print job.
type JobPauseRequest struct {
	JobId int
}

// Do sends an API request and returns an error if any.
func (cmd *JobPauseRequest) Do(c *Client) error {
	uri := fmt.Sprintf("%s/%d/pause", JobApiUri, cmd.JobId)
	_, err := c.doJsonRequest("PUT", uri, nil, JobToolErrors, true)
	return err
}
