package utils

import (
	// "fmt"
	// "sort"
	// "strings"

	// "github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	// "github.com/the-ress/prusalink-screen/pkg/uiWidgets"
)

type FileResponsesSortedByName []*dataModels.FileResponse

func (this FileResponsesSortedByName) Len() int {
	return len(this)
}

func (this FileResponsesSortedByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this FileResponsesSortedByName) Less(i, j int) bool {
	return this[j].Name > this[i].Name
}
