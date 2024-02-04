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
	imageLoader *uiUtils.ImageLoader,
) *gtk.Box {
	infoAndActionRow := uiUtils.MustBox(gtk.ORIENTATION_HORIZONTAL, 5)

	isFolder := fileResponse.IsFolder()

	// Column 1: Folder or File icon.
	var imageFileName uiUtils.ImageFileName
	if isFolder {
		imageFileName = uiUtils.FolderSvg
	} else {
		imageFileName = uiUtils.FileGcodeSvg
	}
	infoAndActionRow.Add(imageLoader.MustGetImageWithSize(imageFileName, fileSystemImageWidth, fileSystemImageHeight))

	// Column 2: File name and file info.
	name := fileResponse.Name
	nameLabel := CreateNameLabel(name)
	infoLabel := CreateInfoLabel(fileResponse, isFolder)
	labelsBox := CreateLabelsBox(nameLabel, infoLabel)
	infoAndActionRow.Add(labelsBox)

	// Column 3: Printer image.
	var actionImage *gtk.Image
	if isFolder {
		actionImage = CreateOpenLocationImage(index, printerImageWidth, printerImageHeight, imageLoader)
	} else {
		actionImage = CreatePrintImage(printerImageWidth, printerImageHeight, imageLoader)
	}

	actionBox := uiUtils.MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	actionBox.Add(actionImage)

	infoAndActionRow.Add(actionBox)

	return infoAndActionRow
}
