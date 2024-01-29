package prusaLinkApis

import (
	"fmt"
)

// JobResumeRequest resumes/resumes/toggles the current print job.
type JobResumeRequest struct {
	JobId int
}

// Do sends an API request and returns an error if any.
func (cmd *JobResumeRequest) Do(c *Client) error {
	uri := fmt.Sprintf("%s/%d/resume", JobApiUri, cmd.JobId)
	_, err := c.doJsonRequest("PUT", uri, nil, JobToolErrors, true)
	return err
}
