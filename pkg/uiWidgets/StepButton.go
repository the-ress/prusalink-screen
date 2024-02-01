package uiWidgets

import (
	"fmt"
	"sync"

	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
)

type Step struct {
	Label         string
	ImageFileName string
	Image         gtk.IWidget
	Value         interface{}
}

type StepButton struct {
	*gtk.Button
	sync.RWMutex

	Steps            []Step
	CurrentStepIndex int
	clicked          func()

	config *config.ScreenConfig
}

func CreateStepButton(
	config *config.ScreenConfig,
	colorVariation int,
	clicked func(),
	steps ...Step,
) *StepButton {
	stepCount := len(steps)
	if stepCount < 1 {
		logger.Error("PANIC!!! - CreateStepButton() - len(steps) < 1")
		panic("StepButton.CreateStepButton() - steps is empty")
	}

	base := uiUtils.MustButtonImageUsingFilePath(config, steps[0].Label, steps[0].ImageFileName, nil)
	if stepCount > 1 {
		ctx, _ := base.GetStyleContext()
		colorClass := fmt.Sprintf("color-dash-%d", colorVariation)
		ctx.AddClass(colorClass)
	}

	instance := &StepButton{
		Button:           base,
		Steps:            steps,
		CurrentStepIndex: -1,
		clicked:          clicked,
		config:           config,
	}

	if stepCount > 0 {
		for i := 0; i < stepCount; i++ {
			instance.Steps[i].Image = uiUtils.MustImageFromFile(config, instance.Steps[i].ImageFileName)
		}

		instance.CurrentStepIndex = 0
	}

	instance.Button.Connect("clicked", instance.handleClick)

	return instance
}

func (this *StepButton) Value() interface{} {
	this.RLock()
	defer this.RUnlock()

	return this.Steps[this.CurrentStepIndex].Value
}

func (this *StepButton) AddStep(step Step) {
	this.Lock()
	defer this.Unlock()

	if this.Steps == nil || len(this.Steps) == 0 {
		this.Steps = make([]Step, 0)
	}

	this.Steps = append(this.Steps, step)
	index := len(this.Steps) - 1
	this.Steps[index].Image = uiUtils.MustImageFromFile(this.config, this.Steps[index].ImageFileName)
}

func (this *StepButton) handleClick() {
	this.RLock()
	defer this.RUnlock()

	stepCount := len(this.Steps)
	if stepCount < 1 {
		return
	}
	this.CurrentStepIndex++
	if this.CurrentStepIndex >= stepCount {
		this.CurrentStepIndex = 0
	}

	this.Button.SetLabel(this.Steps[this.CurrentStepIndex].Label)
	this.Button.SetImage(this.Steps[this.CurrentStepIndex].Image)

	if this.clicked != nil {
		this.clicked()
	}
}
