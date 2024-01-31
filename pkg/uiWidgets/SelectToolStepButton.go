package uiWidgets

import (
	// "fmt"
	"strconv"
	"strings"
	// "github.com/the-ress/prusalink-screen/pkg/logger"
	// "github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	// "github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	// "github.com/the-ress/prusalink-screen/pkg/utils"
)

type SelectToolStepButton struct {
	*StepButton
}

func (this *SelectToolStepButton) Value() string {
	return this.StepButton.Value().(string)
}

func (this *SelectToolStepButton) Index() int {
	value := strings.Replace(this.Value(), "tool", "", -1)
	index, _ := strconv.Atoi(value)

	return index
}
