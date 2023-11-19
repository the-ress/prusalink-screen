package octoprintApis

// "bytes"
// "encoding/json"
// "io"

// "github.com/the-ress/prusalink-screen/octoprintApis/dataModels"

const ConnectionApiUri = "/api/connection"

var ConnectionErrors = StatusMapping{
	400: "The selected port or baudrate for a connect command are not part of the available option",
}
