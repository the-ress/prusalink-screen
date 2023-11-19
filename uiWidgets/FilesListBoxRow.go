package uiWidgets

import (
	// "fmt"

	"github.com/gotk3/gotk3/gtk"

	// "github.com/the-ress/prusalink-screen/logger"
	// "github.com/the-ress/prusalink-screen/octoprintApis"
	"github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
	"github.com/the-ress/prusalink-screen/utils"
	// "github.com/the-ress/prusalink-screen/utils"
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
	pixbufCache *utils.PixbufCache,
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
		pixbufCache,
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
