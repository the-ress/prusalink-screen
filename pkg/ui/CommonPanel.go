package ui

import (
	"fmt"
	"math"

	// "strings"
	// "sync"
	// "time"

	// "github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	// "github.com/the-ress/prusalink-screen/pkg/interfaces"

	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type CommonPanel struct {
	name string
	UI   *UI
	// parentPanel			interfaces.IPanel
	includeBackButton bool
	grid              *gtk.Grid
	preShowCallback   func()
	///backgroundTask		*utils.BackgroundTask
	panelWidth  int
	panelHeight int
	backButton  *gtk.Button
	buttons     []gtk.IWidget
}

func CreateCommonPanel(
	name string,
	ui *UI,
	//parentPanel interfaces.IPanel,
) CommonPanel {
	return newCommonPanel(
		name,
		ui,
		//parentPanel,
		true,
	)
}

func CreateTopLevelCommonPanel(
	name string,
	ui *UI,
	//parentPanel interfaces.IPanel,
) CommonPanel {
	return newCommonPanel(
		name,
		ui,
		//parentPanel,
		false,
	)
}

func newCommonPanel(
	name string,
	ui *UI,
	// parentPanel interfaces.IPanel,
	includeBackButton bool,
) CommonPanel {
	grid := uiUtils.MustGrid()
	grid.SetRowHomogeneous(true)
	grid.SetColumnHomogeneous(true)

	return CommonPanel{
		name: name,
		UI:   ui,
		// parentPanel:		parentPanel,
		includeBackButton: includeBackButton,
		grid:              grid,
		// preShowCallback:
		panelWidth:  4,
		panelHeight: 3,
		// buttons:
	}
}

func (this *CommonPanel) Initialize() {
	last := this.panelWidth * this.panelHeight
	if last < len(this.buttons) {
		cols := math.Ceil(float64(len(this.buttons)) / float64(this.panelWidth))
		last = int(cols) * this.panelWidth
	}

	for i := len(this.buttons) + 1; i < last; i++ {
		box := uiUtils.MustBox(gtk.ORIENTATION_HORIZONTAL, 0)
		this.AddButton(box)
	}

	backImage := this.UI.ImageLoader.MustGetImage(uiUtils.BackSvg)
	this.backButton = uiUtils.MustButtonImage("Back", backImage, this.UI.GoToPreviousPanel)
	if this.includeBackButton {
		this.AddButton(this.backButton)
	}
}

func (this *CommonPanel) AddButton(button gtk.IWidget) {
	x := len(this.buttons) % this.panelWidth
	y := len(this.buttons) / this.panelWidth
	this.grid.Attach(button, x, y, 1, 1)
	this.buttons = append(this.buttons, button)
}

// Begin IPanel implementation
func (this *CommonPanel) Name() string {
	return this.name
}

func (this *CommonPanel) Grid() *gtk.Grid {
	return this.grid
}

func (this *CommonPanel) PreShow() {
	if this.preShowCallback != nil {
		this.preShowCallback()
	}
}

func (this *CommonPanel) Show() {
	/**
	if this.backgroundTask != nil {
		this.backgroundTask.Start()
	}
	**/
}

func (this *CommonPanel) Hide() {
	/**
	if this.backgroundTask != nil {
		this.backgroundTask.Close()
	}
	**/
}

// End IPanel implementation

func (this *CommonPanel) Scaled(s int) int {
	return s * this.UI.scaleFactor
}

func (this *CommonPanel) arrangeMenuItems(
	grid *gtk.Grid,
	items []dataModels.MenuItem,
	cols int,
) {
	for i, item := range items {
		panel := getPanel(this.UI, this, item)
		if panel != nil {
			color := fmt.Sprintf("color%d", (i%4)+1)
			imageFileName := fmt.Sprintf("%s.svg", item.Icon)
			image, err := this.UI.ImageLoader.GetImage(uiUtils.ImageFileName(imageFileName))
			if err != nil {
				panic(err)
			}
			button := uiUtils.MustButtonImageStyle(image, item.Name, color, func() {
				this.UI.GoToPanel(panel)
			})
			grid.Attach(button, (i % cols), i/cols, 1, 1)
		}
	}
}
