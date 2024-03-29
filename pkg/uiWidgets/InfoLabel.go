package uiWidgets

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/gotk3/gotk3/gtk"

	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

func CreateInfoLabel(
	fileResponse *dataModels.FileResponse,
	isFolder bool,
) *gtk.Label {
	label := uiUtils.MustLabel("")
	label.SetHAlign(gtk.ALIGN_START)

	uploadedSize := humanize.Bytes(uint64(fileResponse.Size))
	markup := ""
	if isFolder {
		markup = fmt.Sprintf("<small>Size: <b>%s</b></small>", uploadedSize)
	} else {
		uploadedTime := humanize.Time(fileResponse.Date.Time)
		markup = fmt.Sprintf("<small>Uploaded: <b>%s</b> - Size: <b>%s</b></small>", uploadedTime, uploadedSize)
	}
	label.SetMarkup(markup)

	return label
}
