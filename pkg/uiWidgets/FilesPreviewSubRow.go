package uiWidgets

import (
	"context"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"

	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
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
	imageLoader *uiUtils.ImageLoader,
	previewSubRow *gtk.Box,
	fileResponse *dataModels.FileResponse,
	client *prusaLinkApis.Client,
	fileSystemImageWidth int,
	fileSystemImageHeight int,
) {
	if fileResponse.Refs.Thumbnail == "" {
		return
	}

	logger.Debugf("FilesPreviewSubRow.createPreviewThumbnail() - fileResponse.Refs.Thumbnail is %s", fileResponse.Refs.Thumbnail)

	imageBuffer, imageFromUrlErr := (&prusaLinkApis.ThumbnailRequest{Path: fileResponse.Refs.Thumbnail}).Do(client)
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

		previewImage, imageFromBufferErr := uiUtils.ImageFromBuffer(imageBuffer)

		if imageFromBufferErr != nil {
			logger.Error("FilesPreviewSubRow.createPreviewThumbnail() - error from ImageFromBuffer:", imageFromBufferErr)
			return
		}
		logger.Debug("FilesPreviewSubRow.createPreviewThumbnail() - no error from ImageFromBuffer, now trying to add it...")

		bottomBox := uiUtils.MustBox(gtk.ORIENTATION_HORIZONTAL, 0)

		// Initially was setting the horizontal alignment with CSS, but different resolutions
		// (eg 800x480 vs 480x320) didn't align correctly, so I added a blank SVG to offset
		// the preview thumbnail image.
		spacerImage := imageLoader.MustGetImageWithSize(uiUtils.BlankSvg, fileSystemImageWidth, fileSystemImageHeight)
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
