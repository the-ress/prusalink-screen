package uiUtils

import (
	"fmt"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

type ImageFileName string

const (
	RefreshSvg             ImageFileName = "refresh.svg"
	RestartSvg             ImageFileName = "restart.svg"
	RebootSvg              ImageFileName = "reboot.svg"
	ShutdownSvg            ImageFileName = "shutdown.svg"
	InfoSvg                ImageFileName = "info.svg"
	NetworkSvg             ImageFileName = "network.svg"
	HeatUpSvg              ImageFileName = "heat-up.svg"
	PrintSvg               ImageFileName = "print.svg"
	MoveSvg                ImageFileName = "move.svg"
	FilamentSpoolSvg       ImageFileName = "filament-spool.svg"
	StopSvg                ImageFileName = "stop.svg"
	PrintingControlSvg     ImageFileName = "printing-control.svg"
	CompleteSvg            ImageFileName = "complete.svg"
	ExtruderExtrudeSvg     ImageFileName = "extruder-extrude.svg"
	ExtruderRetractSvg     ImageFileName = "extruder-retract.svg"
	HomeSvg                ImageFileName = "home.svg"
	HomeXSvg               ImageFileName = "home-x.svg"
	HomeYSvg               ImageFileName = "home-y.svg"
	HomeZSvg               ImageFileName = "home-z.svg"
	IncreaseSvg            ImageFileName = "increase.svg"
	DecreaseSvg            ImageFileName = "decrease.svg"
	BackSvg                ImageFileName = "back.svg"
	BlankSvg               ImageFileName = "blank.svg"
	PrusaLinkStorageSvg    ImageFileName = "prusa-link-storage.svg"
	SdSvg                  ImageFileName = "sd.svg"
	BackspaceSvg           ImageFileName = "backspace.svg"
	ToolheadSvg            ImageFileName = "toolhead.svg"
	ExtruderSvg            ImageFileName = "extruder-typeB.svg"
	HotendSvg              ImageFileName = "hotend.svg"
	NozzleSvg              ImageFileName = "nozzle.svg"
	MoveStepSvg            ImageFileName = "move-step.svg"
	BedSvg                 ImageFileName = "bed.svg"
	CoolDownSvg            ImageFileName = "cool-down.svg"
	MoveXPlusSvg           ImageFileName = "move-x+.svg"
	MoveXMinusSvg          ImageFileName = "move-x-.svg"
	MoveYPlusSvg           ImageFileName = "move-y+.svg"
	MoveYMinusSvg          ImageFileName = "move-y-.svg"
	MoveZPlusSvg           ImageFileName = "move-z+.svg"
	MoveZMinusSvg          ImageFileName = "move-z-.svg"
	PauseSvg               ImageFileName = "pause.svg"
	ResumeSvg              ImageFileName = "resume.svg"
	TimeSvg                ImageFileName = "time.svg"
	OpenSvg                ImageFileName = "open.svg"
	FolderSvg              ImageFileName = "folder.svg"
	FileGcodeSvg           ImageFileName = "file-gcode.svg"
	SpeedNormalSvg         ImageFileName = "speed-normal.svg"
	SpeedFastSvg           ImageFileName = "speed-fast.svg"
	SpeedSlowSvg           ImageFileName = "speed-slow.svg"
	FilamentSpoolLoadSvg   ImageFileName = "filament-spool-load.svg"
	FilamentSpoolUnloadSvg ImageFileName = "filament-spool-unload.svg"

	// Logos
	PrusaPrinterLogoPng    ImageFileName = "logos/prusa-printer.png"
	PrusaLinkLogoSvg       ImageFileName = "logos/prusa-link.svg"
	PrusaLinkScreenLogoSvg ImageFileName = "logos/octoscreen-logo.svg"          // TODO
	PrusaLinkScreenLogoPng ImageFileName = "logos/octoscreen-isometric-90%.png" // TODO
	PrusaConnectLogoPng    ImageFileName = "logos/prusa-connect.png"
	SplashLogoPng          ImageFileName = "logos/logo.png"
)

type ImageLoader struct {
	config      *config.ScreenConfig
	pixbufCache *PixbufCache
}

func NewImageLoader(config *config.ScreenConfig) *ImageLoader {
	return &ImageLoader{
		config:      config,
		pixbufCache: NewPixbufCache(),
	}
}

func (this *ImageLoader) GetImage(fileName ImageFileName) (*gtk.Image, error) {
	pixbuf, err := this.GetPixbuf(fileName)
	if err != nil {
		return nil, err
	}

	return this.GetImageFromPixbuf(pixbuf)
}

func (this *ImageLoader) GetImageWithSize(fileName ImageFileName, width, height int) (*gtk.Image, error) {
	pixbuf, err := this.GetPixbufWithSize(fileName, width, height)
	if err != nil {
		return nil, err
	}

	return this.GetImageFromPixbuf(pixbuf)
}

func (this *ImageLoader) GetPixbuf(fileName ImageFileName) (*gdk.Pixbuf, error) {
	pixbuf := this.pixbufCache.GetPixbuf(fileName)
	if pixbuf != nil {
		// Return cached pixbuf
		return pixbuf, nil
	}

	filePath := imagePath(this.config.CssStyleFilePath, string(fileName))
	if !utils.FileExists(filePath) {
		return nil, fmt.Errorf("Image file doesn't exist: %s", filePath)
	}

	pixbuf, err := gdk.PixbufNewFromFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to load image %s: %q", filePath, err)
	}

	this.pixbufCache.SetPixbuf(fileName, pixbuf)

	return pixbuf, nil
}

func (this *ImageLoader) GetPixbufWithSize(fileName ImageFileName, width, height int) (*gdk.Pixbuf, error) {
	pixbuf := this.pixbufCache.GetPixbufWithSize(fileName, width, height)
	if pixbuf != nil {
		// Return cached pixbuf
		return pixbuf, nil
	}

	filePath := imagePath(this.config.CssStyleFilePath, string(fileName))
	if !utils.FileExists(filePath) {
		return nil, fmt.Errorf("Image file doesn't exist: %s", filePath)
	}

	pixbuf, err := gdk.PixbufNewFromFileAtScale(filePath, width, height, true)
	if err != nil {
		return nil, fmt.Errorf("Failed to load image %s: %q", filePath, err)
	}

	this.pixbufCache.SetPixbufWithSize(fileName, width, height, pixbuf)

	return pixbuf, nil
}

func (this *ImageLoader) GetImageFromPixbuf(pixbuf *gdk.Pixbuf) (*gtk.Image, error) {
	image, err := gtk.ImageNewFromPixbuf(pixbuf)
	if err != nil {
		return nil, fmt.Errorf("Failed to create image from pixbuf: %q", err)
	}

	return image, nil
}

func (this *ImageLoader) MustGetImage(fileName ImageFileName) *gtk.Image {
	image, err := this.GetImage(fileName)
	if err != nil {
		panic(err)
	}
	return image
}

func (this *ImageLoader) MustGetImageWithSize(fileName ImageFileName, width, height int) *gtk.Image {
	image, err := this.GetImageWithSize(fileName, width, height)
	if err != nil {
		panic(err)
	}
	return image
}
