package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type FilesListBoxRow struct {
	ClickableListBoxRow

	rowIndex int
}

func CreateFilesListBoxRow(
	fileResponse *dataModels.FileResponse,
	fileSystemImageWidth int,
	fileSystemImageHeight int,
	printerImageWidth int,
	printerImageHeight int,
	rowIndex int,
	rowClickHandler func(button *gtk.Button, index int),
	imageLoader *uiUtils.ImageLoader,
) (*FilesListBoxRow, *gtk.Box) {
	const ROW_PADDING = 0
	base := CreateClickableListBoxRow(rowIndex, ROW_PADDING, rowClickHandler)

	instance := &FilesListBoxRow{
		ClickableListBoxRow: *base,
	}

	styleContext, _ := instance.GetStyleContext()
	styleContext.AddClass("list-item-button")

	isFolder := fileResponse.IsFolder()

	verticalLayoutBox := CreateVerticalLayoutBox()

	// Part I
	filesInfoAndActionSubRow := CreateFilesInfoAndActionSubRow(
		fileResponse,
		rowIndex,
		fileSystemImageWidth,
		fileSystemImageHeight,
		printerImageWidth,
		printerImageHeight,
		imageLoader,
	)
	verticalLayoutBox.Add(filesInfoAndActionSubRow)

	// Part II
	var filesPreviewSubRow *gtk.Box = nil
	if !isFolder {
		filesPreviewSubRow = CreateFilesPreviewSubRow(
			fileResponse,
			fileSystemImageWidth,
			fileSystemImageHeight,
		)
		verticalLayoutBox.Add(filesPreviewSubRow)
	}

	instance.Add(verticalLayoutBox)

	return instance, filesPreviewSubRow
}
