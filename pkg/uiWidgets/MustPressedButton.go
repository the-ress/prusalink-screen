package uiWidgets

import (
	"sync"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/logger"
)

func MustPressedButton(
	label string,
	image *gtk.Image,
	pressed func(),
	speed time.Duration,
) *gtk.Button {
	isReleased := make(chan bool)
	var mutex sync.Mutex

	button, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		logger.LogError("PANIC!!! - MustPressedButton()", "gtk.ButtonNewWithLabel()", err)
		panic(err)
	}

	button.SetImage(image)
	button.SetAlwaysShowImage(true)
	button.SetImagePosition(gtk.POS_TOP)
	button.SetVExpand(true)
	button.SetHExpand(true)

	if pressed != nil {
		button.Connect("pressed", func() {
			go func() {
				for {
					select {
					case <-isReleased:
						return

					default:
						mutex.Lock()
						pressed()
						time.Sleep(speed * time.Millisecond)
						mutex.Unlock()
					}
				}
			}()
		})
	}

	if isReleased != nil {
		button.Connect("released", func() {
			isReleased <- true
		})
	}

	return button
}
