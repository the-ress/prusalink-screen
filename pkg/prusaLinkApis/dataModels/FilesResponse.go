package dataModels

// FilesResponse is the response to a FilesRequest.
type FilesResponse struct {
	// Files is the list of requested files.  Might be an empty list if no files are available
	Files []*FileResponse `json:"children"`
}
