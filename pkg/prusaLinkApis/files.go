package prusaLinkApis

const FilesApiUri = "/api/v1/files"

var (
	FilesLocationGETErrors = StatusMapping{
		404: "Location is neither local nor sdcard",
	}

	FilesLocationPathPOSTErrors = StatusMapping{
		404: "File you want was not found.",
		409: "Printer is not in state to print.",
	}
)
