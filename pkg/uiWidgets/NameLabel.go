package uiWidgets

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

func CreateNameLabel(name string) *gtk.Label {
	label := uiUtils.MustLabel(name)
	markup := fmt.Sprintf("<big>%s</big>", name)
	label.SetMarkup(markup)
	label.SetHExpand(true)
	label.SetEllipsize(pango.ELLIPSIZE_END)
	label.SetHAlign(gtk.ALIGN_START)

	return label
}
