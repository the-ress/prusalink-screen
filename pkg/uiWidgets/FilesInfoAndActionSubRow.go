package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

func CreateFilesInfoAndActionSubRow(
	fileResponse *dataModels.FileResponse,
	index int,
	fileSystemImageWidth int,
	fileSystemImageHeight int,
	printerImageWidth int,
	printerImageHeight int,
	pixbufCache *uiUtils.PixbufCache,
) *gtk.Box {
	infoAndActionRow := uiUtils.MustBox(gtk.ORIENTATION_HORIZONTAL, 5)

	isFolder := fileResponse.IsFolder()

	// Column 1: Folder or File icon.
	var itemImage *gtk.Image
	if isFolder {
		itemImage = pixbufCache.MustImageFromFileWithSize("folder.svg", fileSystemImageWidth, fileSystemImageHeight)
	} else {
		itemImage = pixbufCache.MustImageFromFileWithSize("file-gcode.svg", fileSystemImageWidth, fileSystemImageHeight)
	}
	infoAndActionRow.Add(itemImage)

	// Column 2: File name and file info.
	name := fileResponse.Name
	nameLabel := CreateNameLabel(name)
	infoLabel := CreateInfoLabel(fileResponse, isFolder)
	labelsBox := CreateLabelsBox(nameLabel, infoLabel)
	infoAndActionRow.Add(labelsBox)

	// Column 3: Printer image.
	var actionImage *gtk.Image
	if isFolder {
		actionImage = CreateOpenLocationImage(index, printerImageWidth, printerImageHeight, pixbufCache)
	} else {
		actionImage = CreatePrintImage(printerImageWidth, printerImageHeight, pixbufCache)
	}

	actionBox := uiUtils.MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	actionBox.Add(actionImage)

	infoAndActionRow.Add(actionBox)

	return infoAndActionRow
}
