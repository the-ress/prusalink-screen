package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"

	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type ActionFooter struct {
	gtk.Box

	refreshButton *gtk.Button
	backButton    *gtk.Button
}

func CreateActionFooter(
	config *config.ScreenConfig,
	buttonWidth int,
	buttonHeight int,
	refreshClicked func(),
	backClicked func(),
) *ActionFooter {
	base := uiUtils.MustBox(gtk.ORIENTATION_HORIZONTAL, 5)

	instance := &ActionFooter{
		Box: *base,
	}

	instance.SetHAlign(gtk.ALIGN_END)
	instance.SetHExpand(true)
	instance.SetMarginTop(5)
	instance.SetMarginBottom(5)
	instance.SetMarginEnd(5)

	instance.refreshButton = instance.createRefreshButton(config, buttonWidth, buttonHeight, refreshClicked)
	instance.Add(instance.refreshButton)

	instance.backButton = instance.createBackButton(config, buttonWidth, buttonHeight, backClicked)
	instance.Add(instance.backButton)

	return instance
}

func (this *ActionFooter) createRefreshButton(config *config.ScreenConfig, buttonWidth int, buttonHeight int, clicked func()) *gtk.Button {
	image := uiUtils.MustImageFromFileWithSize(config, "refresh.svg", buttonWidth, buttonHeight)
	return uiUtils.MustButton(image, clicked)
}

func (this *ActionFooter) createBackButton(config *config.ScreenConfig, buttonWidth int, buttonHeight int, clicked func()) *gtk.Button {
	image := uiUtils.MustImageFromFileWithSize(config, "back.svg", buttonWidth, buttonHeight)
	return uiUtils.MustButton(image, clicked)
}
