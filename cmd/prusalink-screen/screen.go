package main

import (
	standardLog "log"
	"runtime"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"

	"github.com/the-ress/prusalink-screen/pkg/config"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/ui"
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

func main() {
	defer mainDefer()

	logger.TraceEnter("PrusaLinkScreen - main.main()")

	startSystemDHeartbeat()

	initializeGtk()

	config, err := config.ReadConfig()
	if err != nil {
		panic(err.Error())
	}

	setLogLevel(config.LogLevel)

	utils.DumpSystemInformation()
	utils.DumpEnvironmentVariables()
	config.Dump()

	setCursor(config.DisplayCursor)

	_ = ui.NewUi(config)

	gtk.Main()

	logger.TraceLeave("OctoScreen - main.main()")
}

func startSystemDHeartbeat() {
	systemDHeartbeat := utils.GetSystemDHeartbeatInstance()
	systemDHeartbeat.Start()
}

func initializeGtk() {
	gtk.Init(nil)
	gtkSettings, _ := gtk.SettingsGetDefault()
	gtkSettings.SetProperty("gtk-application-prefer-dark-theme", true)
}

func setLogLevel(logLevel string) {
	switch logLevel {
	case "debug":
		logger.SetLogLevel(logrus.DebugLevel)

	case "info":
		logger.SetLogLevel(logrus.InfoLevel)

	case "warn":
		logger.SetLogLevel(logrus.WarnLevel)

	case "error":
		logger.SetLogLevel(logrus.ErrorLevel)

	default:
		logger.Errorf("main.setLogLevel() - unknown logLevel: %q, defaulting to error", logLevel)
		logLevel = "error"
		logger.SetLogLevel(logrus.ErrorLevel)
	}

	standardLog.Printf("main.SetLogLevel() - logLevel is now set to: %q", logLevel)
}

func setCursor(displayCursor bool) {
	// For reference, see "How to turn on a pointer"
	// (https://github.com/Z-Bolt/OctoScreen/issues/285)
	// ...and "No mouse pointer when running xinit"
	// (https://www.raspberrypi.org/forums/viewtopic.php?t=139546)

	if !displayCursor {
		return
	}

	window, err := getRootWindow()
	if err != nil {
		return
	}

	cursor, err := getDefaultCursor()
	if err != nil {
		return
	}

	window.SetCursor(cursor)
}

func getRootWindow() (*gdk.Window, error) {
	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		return nil, err
	}

	window, err := screen.GetRootWindow()

	return window, err
}

func getDefaultCursor() (*gdk.Cursor, error) {
	display, err := gdk.DisplayGetDefault()
	if err != nil {
		return nil, err
	}

	// Examples of the different cursors can be found at
	// https://developer.gnome.org/gdk3/stable/gdk3-Cursors.html#gdk-cursor-new-from-name
	cursor, err := gdk.CursorNewFromName(display, "default")

	return cursor, err
}

func mainDefer() {
	standardLog.Println("main's defer() was called, now calling recover()")
	rec := recover()
	if rec != nil {
		standardLog.Println("main's defer() - recover:", rec)
	} else {
		standardLog.Println("main's defer() - recover was nil")
	}

	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	/*
		programCounter, fileName, lineNumber, infoWasRecovered := runtime.Caller(2)
		standardLog.Println("main's defer() - programCounter:", programCounter)
		standardLog.Println("main's defer() - fileName:", fileName)
		standardLog.Println("main's defer() - lineNumber:", lineNumber)
		standardLog.Println("main's defer() - infoWasRecovered:", infoWasRecovered)
	*/

	pc := make([]uintptr, 20)
	numberOfPcEntries := runtime.Callers(0, pc)
	if numberOfPcEntries > 10 {
		numberOfPcEntries = 10
	}

	for i := 1; i < numberOfPcEntries; i++ {
		/*
			standardLog.Printf("main's defer() - [%d]", i)
			standardLog.Printf("main's defer() - [%d]", numberOfPcEntries)

			programCounter, fileName, lineNumber, infoWasRecovered := runtime.Caller(i)
			standardLog.Printf("main's defer() - programCounter[%d]: %v", i, programCounter)
			standardLog.Printf("main's defer() - fileName[%d]: %v", i, fileName)
			standardLog.Printf("main's defer() - lineNumber[%d]: %v", i, lineNumber)
			standardLog.Printf("main's defer() - infoWasRecovered[%d]: %v", i, infoWasRecovered)
			standardLog.Println("")
		*/

		_, fileName, lineNumber, infoWasRecovered := runtime.Caller(i)
		if infoWasRecovered {
			standardLog.Printf("main's defer() - [%d] %s, line %d", i, fileName, lineNumber)
		}
	}

	standardLog.Println("main's defer() was called, now exiting func()")
}
