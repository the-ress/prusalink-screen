package uiWidgets

import (
	// "fmt"

	"github.com/gotk3/gotk3/gtk"

	// "github.com/the-ress/prusalink-screen/octoprintApis/dataModels"
	"github.com/the-ress/prusalink-screen/utils"
)

type SystemInfoBox struct {
	*gtk.Box

	titleLabel *gtk.Label
	textLabel  *gtk.Label
}

func NewSystemInfoBox(
	image *gtk.Image,
	title string,
) *SystemInfoBox {
	base := utils.MustBox(gtk.ORIENTATION_VERTICAL, 0)
	base.SetHExpand(true)
	base.SetHAlign(gtk.ALIGN_CENTER)
	base.SetVExpand(true)
	base.SetVAlign(gtk.ALIGN_CENTER)

	ctx, _ := image.GetStyleContext()
	ctx.AddClass("margin-top-5")
	base.Add(image)

	titleLabel := utils.MustLabel(title)
	ctx, _ = titleLabel.GetStyleContext()
	ctx.AddClass("margin-top-10")
	ctx.AddClass("font-size-18")
	base.Add(titleLabel)

	textLabel := utils.MustLabel("")
	ctx, _ = textLabel.GetStyleContext()
	ctx.AddClass("font-size-18")
	base.Add(textLabel)

	return &SystemInfoBox{
		Box:        base,
		titleLabel: titleLabel,
		textLabel:  textLabel,
	}
}

func (this *SystemInfoBox) SetText(text string) {
	this.textLabel.SetMarkup(text)
}
