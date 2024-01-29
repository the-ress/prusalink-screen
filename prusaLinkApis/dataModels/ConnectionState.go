package dataModels

import (
	"strings"
)

type ConnectionState string

func (s ConnectionState) IsError() bool {
	return strings.HasPrefix(string(s), "Error")
}
