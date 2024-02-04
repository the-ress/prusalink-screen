package prusaLinkApis

import (
	"encoding/json"

	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
)

// Gets the list of storage locations
type StorageRequest struct{}

func (cmd *StorageRequest) Do(c *Client) (*dataModels.StorageResponse, error) {
	bytes, err := c.doJsonRequest("GET", StorageApiUri, nil, nil, true)
	if err != nil {
		return nil, err
	}

	response := &dataModels.StorageResponse{}
	if err := json.Unmarshal(bytes, response); err != nil {
		return nil, err
	}

	return response, err
}
