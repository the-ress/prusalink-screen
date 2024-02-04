package dataModels

type StorageResponse struct {
	StorageList []StorageLocation `json:"storage_list"`
}

type StorageLocation struct {
	Type      string `json:"type"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
}
