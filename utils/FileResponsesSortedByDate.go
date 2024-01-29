package utils

import (
	// "fmt"
	// "sort"

	// "github.com/the-ress/prusalink-screen/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/prusaLinkApis/dataModels"
	// "github.com/the-ress/prusalink-screen/uiWidgets"
)

type FileResponsesSortedByDate []*dataModels.FileResponse

func (this FileResponsesSortedByDate) Len() int {
	return len(this)
}

func (this FileResponsesSortedByDate) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this FileResponsesSortedByDate) Less(i, j int) bool {
	return this[j].Date.Time.Before(this[i].Date.Time)
}
