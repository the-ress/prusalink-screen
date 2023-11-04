package octoprintApis

type ThumbnailRequest struct {
	Path string
}

// Do sends an API request and returns the API response.
func (cmd *ThumbnailRequest) Do(c *Client) ([]byte, error) {
	return c.doRequest("GET", cmd.Path, "", nil, nil, true)
}
