package uiUtils

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/common"
	"github.com/the-ress/prusalink-screen/pkg/logger"
)

// MustWindow returns a new gtk.Window, if error panics.
func MustWindow(windowType gtk.WindowType) *gtk.Window {
	win, err := gtk.WindowNew(windowType)
	if err != nil {
		logger.LogError("PANIC!!! - MustWindow()", "gtk.WindowNew()", err)
		panic(err)
	}

	win.SetResizable(false)

	return win
}

// MustGrid returns a new gtk.Grid, if error panics.
func MustGrid() *gtk.Grid {
	grid, err := gtk.GridNew()
	if err != nil {
		logger.LogError("PANIC!!! - MustGrid()", "gtk.GridNew()", err)
		panic(err)
	}

	return grid
}

// MustBox returns a new gtk.Box, with the given configuration, if err panics.
func MustBox(orientation gtk.Orientation, spacing int) *gtk.Box {
	box, err := gtk.BoxNew(orientation, spacing)
	if err != nil {
		logger.LogError("PANIC!!! - MustBox()", "gtk.BoxNew()", err)
		panic(err)
	}

	return box
}

// MustProgressBar returns a new gtk.ProgressBar, if err panics.
func MustProgressBar() *gtk.ProgressBar {
	progressBar, err := gtk.ProgressBarNew()
	if err != nil {
		logger.LogError("PANIC!!! - MustProgressBar()", "gtk.ProgressBarNew()", err)
		panic(err)
	}

	return progressBar
}

// MustLabel returns a new gtk.Label, if err panics.
func MustLabel(format string, args ...interface{}) *gtk.Label {
	label, err := gtk.LabelNew("")
	if err != nil {
		logger.LogError("PANIC!!! - MustLabel()", "gtk.LabelNew()", err)
		panic(err)
	}

	label.SetMarkup(fmt.Sprintf(format, args...))

	return label
}

// MustLabelWithCssClass returns a stylized new gtk.Label, if err panics.
func MustLabelWithCssClass(format string, className string, args ...interface{}) *gtk.Label {
	label, err := gtk.LabelNew("")
	if err != nil {
		logger.LogError("PANIC!!! - MustLabelWithCssClass()", "gtk.LabelNew()", err)
		panic(err)
	}

	ctx, _ := label.GetStyleContext()
	ctx.AddClass(className)

	label.SetMarkup(fmt.Sprintf(format, args...))

	return label
}

// MustLabelWithCssClass returns a stylized new gtk.Label, if err panics.
func MustLabelWithCssClasses(format string, classNames []string, args ...interface{}) *gtk.Label {
	label, err := gtk.LabelNew("")
	if err != nil {
		logger.LogError("PANIC!!! - MustLabelWithCssClasses()", "gtk.LabelNew()", err)
		panic(err)
	}

	label.SetMarkup(fmt.Sprintf(format, args...))

	ctx, _ := label.GetStyleContext()
	for i := 0; i < len(classNames); i++ {
		ctx.AddClass(classNames[i])
	}

	return label
}

// LabelWithImage represents a gtk.Label with a image to the right.
type LabelWithImage struct {
	Label *gtk.Label
	*gtk.Box
}

// LabelImageSize default width and height of the image for a LabelWithImage
const LabelImageSize = 20

// MustLabelWithImage returns a new LabelWithImage based on a gtk.Box containing
// a gtk.Label with a gtk.Image, the image is scaled at LabelImageSize.
func MustLabelWithImage(imageLoader *ImageLoader, imageFileName ImageFileName, format string, args ...interface{}) *LabelWithImage {
	label := MustLabel(format, args...)
	box := MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	box.Add(imageLoader.MustGetImageWithSize(imageFileName, LabelImageSize, LabelImageSize))
	box.Add(label)

	return &LabelWithImage{Label: label, Box: box}
}

// MustButtonImageStyle returns a new gtk.Button with the given label, image and clicked callback, if error panics.
func MustButtonImageStyle(image *gtk.Image, buttonLabel, style string, clicked func()) *gtk.Button {
	button := MustButtonImage(buttonLabel, image, clicked)
	ctx, _ := button.GetStyleContext()
	ctx.AddClass(style)

	return button
}
func MustButtonImage(buttonLabel string, image *gtk.Image, clicked func()) *gtk.Button {
	button, err := gtk.ButtonNewWithLabel(buttonLabel)
	if err != nil {
		logger.LogError("PANIC!!! - MustButtonImage()", "gtk.ButtonNewWithLabel()", err)
		panic(err)
	}

	button.SetImage(image)
	button.SetAlwaysShowImage(true)
	button.SetImagePosition(gtk.POS_TOP)
	button.SetVExpand(true)
	button.SetHExpand(true)

	if clicked != nil {
		button.Connect("clicked", clicked)
	}

	return button
}

func MustToggleButton(label string, image *gtk.Image, clicked func()) *gtk.ToggleButton {
	button, err := gtk.ToggleButtonNewWithLabel(label)
	if err != nil {
		logger.LogError("PANIC!!! - MustToggleButton()", "gtk.ToggleButtonNewWithLabel()", err)
		panic(err)
	}

	button.SetImage(image)
	button.SetAlwaysShowImage(true)
	button.SetImagePosition(gtk.POS_TOP)
	button.SetVExpand(true)
	button.SetHExpand(true)

	if clicked != nil {
		button.Connect("clicked", clicked)
	}

	return button
}

func MustButton(image *gtk.Image, clicked func()) *gtk.Button {
	button, err := gtk.ButtonNew()
	if err != nil {
		logger.LogError("PANIC!!! - MustButton()", "gtk.ButtonNew()", err)
		panic(err)
	}

	button.SetImage(image)
	button.SetImagePosition(gtk.POS_TOP)

	if clicked != nil {
		button.Connect("clicked", clicked)
	}

	return button
}

func MustButtonText(label string, clicked func()) *gtk.Button {
	button, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		logger.LogError("PANIC!!! - MustButtonText()", "gtk.ButtonNewWithLabel()", err)
		panic(err)
	}

	if clicked != nil {
		button.Connect("clicked", clicked)
	}

	return button
}

func ImageFromBuffer(buffer []byte) (*gtk.Image, error) {
	pixbufLoader, err := gdk.PixbufLoaderNew()
	if err != nil {
		return nil, err
	}
	defer pixbufLoader.Close()

	writeLength, err := pixbufLoader.Write(buffer)
	if err != nil {
		return nil, err
	} else if writeLength < 1 {
		return nil, errors.New("bytes written was zero")
	}

	pixbuf, err := pixbufLoader.GetPixbuf()
	if err != nil {
		return nil, err
	}

	image, err := gtk.ImageNewFromPixbuf(pixbuf)

	return image, err
}

func ImageFromBufferAtSize(buffer []byte, width int, height int) (*gtk.Image, error) {
	pixbufLoader, err := gdk.PixbufLoaderNew()
	if err != nil {
		return nil, err
	}
	defer pixbufLoader.Close()

	writeLength, err := pixbufLoader.Write(buffer)
	if err != nil {
		return nil, err
	} else if writeLength < 1 {
		return nil, errors.New("bytes written was zero")
	}

	pixbuf, err := pixbufLoader.GetPixbuf()
	if err != nil {
		return nil, err
	}

	scaledWidth, scaledHeight := GetScaledSize(
		float64(pixbuf.GetWidth()),
		float64(pixbuf.GetHeight()),
		float64(width),
		float64(height),
	)

	pixbuf, err = pixbuf.ScaleSimple(int(scaledWidth), int(scaledHeight), gdk.INTERP_BILINEAR)
	if err != nil {
		return nil, err
	}
	image, err := gtk.ImageNewFromPixbuf(pixbuf)

	return image, err
}

func GetScaledSize(originalWidth, originalHeight, destWidth, destHeight float64) (float64, float64) {
	if originalWidth == 0 || originalHeight == 0 {
		return originalWidth, originalHeight
	}

	widthRatio := destWidth / originalWidth
	heightRatio := destHeight / originalHeight

	if heightRatio < widthRatio {
		width := originalWidth * heightRatio
		height := originalHeight * heightRatio
		return width, height
	} else {
		width := originalWidth * widthRatio
		height := originalHeight * widthRatio
		return width, height
	}

}

// MustCSSProviderFromFile returns a new gtk.CssProvider for a given css file, if error panics.
func MustCssProviderFromFile(cssStyleFilePath string, cssFileName string) *gtk.CssProvider {
	cssProvider, err := gtk.CssProviderNew()
	if err != nil {
		logger.LogError("PANIC!!! - MustCssProviderFromFile()", "gtk.CssProviderNew()", err)
		panic(err)
	}

	cssFilePath := cssFilePath(cssStyleFilePath, cssFileName)
	if err := cssProvider.LoadFromPath(cssFilePath); err != nil {
		logger.LogError("PANIC!!! - MustCssProviderFromFile()", "cssProvider.LoadFromPath()", err)
		panic(err)
	}

	return cssProvider
}

func cssFilePath(cssStyleFilePath string, cssFileName string) string {
	return filepath.Join(cssStyleFilePath, cssFileName)
}

func imagePath(cssStyleFilePath string, fileName string) string {
	return filepath.Join(cssStyleFilePath, common.ImageFolder, fileName)
}

// MustOverlay returns a new gtk.Overlay, if error panics.
func MustOverlay() *gtk.Overlay {
	overlay, err := gtk.OverlayNew()
	if err != nil {
		logger.LogError("PANIC!!! - MustOverlay()", "gtk.OverlayNew()", err)
		panic(err)
	}

	return overlay
}
