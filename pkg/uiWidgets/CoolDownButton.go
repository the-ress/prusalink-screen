package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"

	// "github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

type CoolDownButton struct {
	*gtk.Button

	client   *prusaLinkApis.Client
	callback func()
}

func CreateCoolDownButton(
	client *prusaLinkApis.Client,
	config *utils.ScreenConfig,
	callback func(),
) *CoolDownButton {
	base := utils.MustButtonImageUsingFilePath(config, "All Off", "cool-down.svg", nil)

	instance := &CoolDownButton{
		Button:   base,
		client:   client,
		callback: callback,
	}
	instance.Button.Connect("clicked", instance.handleClicked)

	return instance
}

func (this *CoolDownButton) handleClicked() {
	TurnAllHeatersOff(this.client)

	if this.callback != nil {
		this.callback()
	}
}

func TurnAllHeatersOff(
	client *prusaLinkApis.Client,
) {
	// Set the bed's temp.
	bedTargetRequest := &prusaLinkApis.BedTargetRequest{
		Target: 0.0,
	}
	err := bedTargetRequest.Do(client)
	if err != nil {
		logger.LogError("CoolDownButton.TurnAllHeatersOff()", "Do(BedTargetRequest)", err)
		return
	}

	// Set the temp of hotend.
	var toolTargetRequest = &prusaLinkApis.ToolTargetRequest{
		Targets: map[string]float64{
			"tool0": 0.0,
		},
	}
	err = toolTargetRequest.Do(client)
	if err != nil {
		logger.LogError("CoolDownButton.TurnAllHeatersOff()", "Do(ToolTargetRequest)", err)
	}

}
