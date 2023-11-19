package uiWidgets

import (
	"context"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/the-ress/prusalink-screen/logger"
	"github.com/the-ress/prusalink-screen/octoprintApis"

	"github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
	"github.com/the-ress/prusalink-screen/utils"
)

func CreateFilesPreviewSubRow(
	fileResponse *dataModels.FileResponse,
	fileSystemImageWidth int,
	fileSystemImageHeight int,
) *gtk.Box {
	return CreateVerticalLayoutBox()
}

func CreatePreviewThumbnail(
	ctx context.Context,
	previewSubRow *gtk.Box,
	fileResponse *dataModels.FileResponse,
	client *octoprintApis.Client,
	fileSystemImageWidth int,
	fileSystemImageHeight int,
) {
	if fileResponse.Refs.Thumbnail == "" {
		return
	}

	logger.Debugf("FilesPreviewSubRow.createPreviewThumbnail() - fileResponse.Refs.Thumbnail is %s", fileResponse.Refs.Thumbnail)

	imageBuffer, imageFromUrlErr := (&octoprintApis.ThumbnailRequest{Path: fileResponse.Refs.Thumbnail}).Do(client)
	if imageFromUrlErr != nil {
		logger.Error("FilesPreviewSubRow.createPreviewThumbnail() - error from ThumbnailRequest:", imageFromUrlErr)
		return
	}

	logger.Debug("FilesPreviewSubRow.createPreviewThumbnail() - no error from ThumbnailRequest, now trying to parse it...")

	glib.IdleAddPriority(glib.PRIORITY_LOW, func() {
		select {
		case <-ctx.Done():
			return // Abort
		default:
		}

		previewImage, imageFromBufferErr := utils.ImageFromBuffer(imageBuffer)

		if imageFromBufferErr != nil {
			logger.Error("FilesPreviewSubRow.createPreviewThumbnail() - error from ImageFromBuffer:", imageFromUrlErr)
			return
		}
		logger.Debug("FilesPreviewSubRow.createPreviewThumbnail() - no error from ImageFromBuffer, now trying to add it...")

		bottomBox := utils.MustBox(gtk.ORIENTATION_HORIZONTAL, 0)

		// Initially was setting the horizontal alignment with CSS, but different resolutions
		// (eg 800x480 vs 480x320) didn't align correctly, so I added a blank SVG to offset
		// the preview thumbnail image.
		spacerImage := utils.MustImageFromFileWithSize("blank.svg", fileSystemImageWidth, fileSystemImageHeight)
		bottomBox.Add(spacerImage)

		// Still need some CSS for the bottom margin.
		previewImageStyleContext, _ := previewImage.GetStyleContext()
		previewImageStyleContext.AddClass("preview-image-list-item")

		// OK, now add the preview image.
		bottomBox.Add(previewImage)

		// ...and finally add everything to the bottom box/container.
		// listItemBox.Add(bottomBox)
		bottomBox.ShowAll()

		previewSubRow.Add(bottomBox)
	})
}
