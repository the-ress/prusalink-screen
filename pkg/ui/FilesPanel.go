package ui

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/gotk3/gotk3/gtk"
	"github.com/the-ress/prusalink-screen/pkg/logger"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis"
	"github.com/the-ress/prusalink-screen/pkg/prusaLinkApis/dataModels"
	"github.com/the-ress/prusalink-screen/pkg/uiUtils"
	"github.com/the-ress/prusalink-screen/pkg/uiWidgets"
	"github.com/the-ress/prusalink-screen/pkg/utils"
)

type filesPanel struct {
	CommonPanel

	scrollableListBox *uiWidgets.ScrollableListBox
	filesListBoxRows  []*uiWidgets.FilesListBoxRow
	actionFooter      *uiWidgets.ActionFooter

	locationHistory      utils.LocationHistory
	currentFileResponses []*dataModels.FileResponse

	pixbufCache             *uiUtils.PixbufCache
	cancelThumbnailDownload context.CancelFunc
}

type filesPanelRow struct {
	model         *dataModels.FileResponse
	previewSubRow *gtk.Box
}

var filesPanelInstance *filesPanel

func GetFilesPanelInstance(
	ui *UI,
) *filesPanel {
	if filesPanelInstance == nil {
		locationHistory := utils.LocationHistory{
			Locations: []dataModels.Location{},
		}

		instance := &filesPanel{
			CommonPanel:     CreateCommonPanel("FilesPanel", ui),
			locationHistory: locationHistory,
		}

		instance.initializeUi()
		instance.initializeData()
		filesPanelInstance = instance
	}

	return filesPanelInstance
}

func (this *filesPanel) initializeUi() {
	this.Grid().SetRowHomogeneous(false)
	this.CreateListBox()
	this.CreateFooter()
}

func (this *filesPanel) CreateListBox() {
	this.scrollableListBox = uiWidgets.CreateScrollableListBox()
	this.Grid().Attach(this.scrollableListBox, 0, 0, 4, 2)
}

func (this *filesPanel) CreateFooter() {
	this.actionFooter = uiWidgets.CreateActionFooter(
		this.UI.ImageLoader,
		this.Scaled(40),
		this.Scaled(40),

		// this.doClear, // just for testing
		this.doLoadFiles,

		this.goBack,
	)
	this.Grid().Attach(this.actionFooter, 2, 2, 2, 1)
}

func (this *filesPanel) initializeData() {
	logger.TraceEnter("FilesPanel.initializeData()")

	this.doLoadFiles()

	logger.TraceLeave("FilesPanel.initializeData()")
}

// doClear() is here just for testing
func (this *filesPanel) doClear() {
	logger.TraceEnter("FilesPanel.doClear()")

	listBoxContainer := this.scrollableListBox.ListBoxContainer()
	uiUtils.EmptyTheContainer(listBoxContainer)

	logger.TraceLeave("FilesPanel.doClear()")
}

func (this *filesPanel) doLoadFiles() {
	logger.TraceEnter("FilesPanel.doLoadFiles()")

	if this.cancelThumbnailDownload != nil {
		this.cancelThumbnailDownload()
		this.cancelThumbnailDownload = nil
	}

	listBoxContainer := this.scrollableListBox.ListBoxContainer()
	uiUtils.EmptyTheContainer(listBoxContainer)

	atRootLevel := this.displayRootLocations()
	// If we are at the "root" level (where the option for Local and SD are displayed),
	// but SD is not ready, push us up into Local so the user doesn't have to work harder than
	// they have to.
	if atRootLevel && !this.sdIsReady() {
		atRootLevel = false
		this.locationHistory = utils.LocationHistory{
			Locations: []dataModels.Location{dataModels.Local},
		}
	}

	if atRootLevel {
		this.addRootLocations()
	} else {
		sortedFiles := this.getSortedFiles()

		/*
			logger.Debugf("FilesPanel.doLoadFiles() sortedFiles:")
			for i := 0; i < len(sortedFiles); i++ {
				sortedFile := sortedFiles[i]
				isFolder := "false"
				if sortedFile.IsFolder() {
					isFolder = "TRUE"
				}
				logger.Debugf("FilesPanel.doLoadFiles() - sortedFiles[%d]:%s, isFolder:%s", i, sortedFile.Name, isFolder)
			}
		*/

		this.addSortedFiles(sortedFiles)
	}

	this.scrollableListBox.ShowAll()

	logger.TraceLeave("FilesPanel.doLoadFiles()")
}

func (this *filesPanel) sdIsReady() bool {
	response, err := (&prusaLinkApis.StorageRequest{}).Do(this.UI.Client)
	if err != nil {
		return false
	}

	for _, v := range response.StorageList {
		if v.Type == "SDCARD" {
			return v.Available
		}
	}

	return false
}

func (this *filesPanel) goBack() {
	if this.displayRootLocations() {
		this.UI.GoToPreviousPanel()
	} else if this.locationHistory.IsRoot() {
		if this.sdIsReady() {
			this.locationHistory.GoBack()
			this.doLoadFiles()
		} else {
			this.UI.GoToPreviousPanel()
		}
	} else {
		this.locationHistory.GoBack()
		this.doLoadFiles()
	}
}

func (this *filesPanel) displayRootLocations() bool {
	if this.locationHistory.Length() < 1 {
		return true
	} else {
		return false
	}
}

func (this *filesPanel) getSortedFiles() []*dataModels.FileResponse {
	var files []*dataModels.FileResponse

	if this.displayRootLocations() {
		return nil
	}

	current := this.locationHistory.CurrentLocation()
	logger.Infof("Loading list of files from: %s", string(current))

	// if current == dataModels.SDCard {
	// 	sdRefreshRequest := &prusaLinkApis.SdRefreshRequest{}
	// 	err := sdRefreshRequest.Do(this.UI.Client)
	// 	if err != nil {
	// 		logger.LogError("getSortedFiles()", "sdRefreshRequest.Do()", err)
	// 		return []*dataModels.FileResponse{}
	// 	} else {
	// 		// Pause here for a second, because the preceding call to filesRequest.Do()
	// 		// doesn't work, and it returns a truncated list of files.  Pausing here
	// 		// for a second seems to resolve the issue.
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }

	filesRequest := &prusaLinkApis.FilesRequest{
		Location:  current,
		Recursive: false,
	}
	filesResponse, err := filesRequest.Do(this.UI.Client)
	if err != nil {
		logger.LogError("files.getSortedFiles()", "Do(FilesRequest)", err)
		files = []*dataModels.FileResponse{}
	} else {
		files = filesResponse.Files
	}

	var filteredFiles []*dataModels.FileResponse
	for i := range files {
		if !strings.HasPrefix(files[i].Path, "trash") {
			filteredFiles = append(filteredFiles, files[i])
		}
	}

	sortedFiles := utils.FileResponsesSortedByDate(filteredFiles)
	// sortedFiles := utils.FileResponsesSortedByName(filteredFiles)
	sort.Sort(sortedFiles)

	return sortedFiles
}

func (this *filesPanel) addRootLocations() {
	this.addMessage("Select source location:")

	response, err := (&prusaLinkApis.StorageRequest{}).Do(this.UI.Client)
	if err != nil {
		return
	}

	for i, v := range response.StorageList {
		if !v.Available {
			continue
		}
		this.addRootLocation(i, v)
	}
}

func (this *filesPanel) addMessage(message string) {
	nameLabel := uiWidgets.CreateNameLabel(message)
	labelsBox := uiWidgets.CreateLabelsBox(nameLabel, nil)
	labelsBox.SetMarginStart(10)

	topBox := uiUtils.MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	topBox.Add(labelsBox)

	listItemBox := uiWidgets.CreateVerticalLayoutBox()
	listItemBox.Add(topBox)

	listItemFrame, _ := gtk.FrameNew("")
	listItemFrame.Add(listItemBox)

	this.scrollableListBox.Add(listItemFrame)
}

func (this *filesPanel) addRootLocation(i int, location dataModels.StorageLocation) {
	rootLocationButton := this.createRootLocationButton(i, location)

	listBoxRow, _ := gtk.ListBoxRowNew()
	listBoxRow.Add(rootLocationButton)

	this.scrollableListBox.Add(listBoxRow)
}

func (this *filesPanel) createRootLocationButton(i int, location dataModels.StorageLocation) *gtk.Button {
	var itemImage *gtk.Image
	if location.Type == "LOCAL" {
		itemImage = this.UI.ImageLoader.MustGetImageWithSize(uiUtils.PrusaLinkStorageSvg, this.Scaled(35), this.Scaled(35))
	} else {
		itemImage = this.UI.ImageLoader.MustGetImageWithSize(uiUtils.SdSvg, this.Scaled(35), this.Scaled(35))
	}

	topBox := uiUtils.MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	topBox.Add(itemImage)

	name := fmt.Sprintf("  %s", location.Name)
	nameLabel := uiWidgets.CreateNameLabel(name)

	infoLabel := uiUtils.MustLabel("")
	infoLabel.SetHAlign(gtk.ALIGN_START)
	infoLabel.SetMarkup("<small> </small>")

	labelsBox := uiWidgets.CreateLabelsBox(nameLabel, infoLabel)
	topBox.Add(labelsBox)

	actionImage := uiWidgets.CreateOpenLocationImage(i, this.Scaled(40), this.Scaled(40), this.UI.ImageLoader)

	actionBox := uiUtils.MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	actionBox.Add(actionImage)
	topBox.Add(actionBox)

	rootLocationButton, _ := gtk.ButtonNew()
	rootLocationButton.Connect("clicked", func() {
		this.locationHistory = utils.LocationHistory{
			Locations: []dataModels.Location{dataModels.Location(location.Path)},
		}

		this.doLoadFiles()
	})

	rootLocationButton.Add(topBox)

	return rootLocationButton
}

func (this *filesPanel) addSortedFiles(sortedFiles []*dataModels.FileResponse) {
	var index int = 0

	rows := make([]filesPanelRow, len(sortedFiles))

	this.currentFileResponses = make([]*dataModels.FileResponse, 0)

	for _, fileResponse := range sortedFiles {
		if fileResponse.IsFolder() {
			this.currentFileResponses = append(this.currentFileResponses, fileResponse)

			filesListBoxRow, _ := uiWidgets.CreateFilesListBoxRow(
				fileResponse,
				this.Scaled(35),
				this.Scaled(35),
				this.Scaled(40),
				this.Scaled(40),
				index,
				this.handleFolderClick,
				this.UI.ImageLoader,
			)
			this.filesListBoxRows = append(this.filesListBoxRows, filesListBoxRow)
			this.scrollableListBox.Add(filesListBoxRow)

			rows[index] = filesPanelRow{
				model: fileResponse,
			}

			index++
		}
	}

	for _, fileResponse := range sortedFiles {
		if !fileResponse.IsFolder() {
			this.currentFileResponses = append(this.currentFileResponses, fileResponse)

			filesListBoxRow, previewSubRow := uiWidgets.CreateFilesListBoxRow(
				fileResponse,
				this.Scaled(35),
				this.Scaled(35),
				this.Scaled(40),
				this.Scaled(40),
				index,
				this.handleFileClick,
				this.UI.ImageLoader,
			)
			this.filesListBoxRows = append(this.filesListBoxRows, filesListBoxRow)
			this.scrollableListBox.Add(filesListBoxRow)

			rows[index] = filesPanelRow{
				model:         fileResponse,
				previewSubRow: previewSubRow,
			}

			index++
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	this.cancelThumbnailDownload = cancel
	go this.downloadThumbnails(ctx, rows)
}

func (this *filesPanel) downloadThumbnails(
	ctx context.Context,
	rows []filesPanelRow,
) {
	logger.TraceEnter("FilesPanel.downloadThumbnails()")

	for _, row := range rows {
		select {
		case <-ctx.Done():
			return // Abort
		default:
		}

		if row.previewSubRow != nil {
			uiWidgets.CreatePreviewThumbnail(
				ctx,
				this.UI.ImageLoader,
				row.previewSubRow,
				row.model,
				this.UI.Client,
				this.Scaled(35),
				this.Scaled(35),
			)
		}
	}

	logger.TraceLeave("FilesPanel.downloadThumbnails()")
}

func (this *filesPanel) handleFolderClick(button *gtk.Button, rowIndex int) {
	logger.TraceEnter("FilesPanel.handleFolderClick()")

	logger.Debugf("FilesPanel.handleFolderClick() - rowIndex: %d", rowIndex)

	if this.currentFileResponses == nil {
		logger.Fatalf("FilesPanel.handleFolderClick() - currentFileResponses is nil")
	}

	fileResponsesLength := len(this.currentFileResponses)
	if rowIndex >= fileResponsesLength {
		logger.Fatalf(
			"FilesPanel.handleFolderClick() - rowIndex is out of range.  The length of currentFileResponses is %d but rowIndex is %d",
			fileResponsesLength,
			rowIndex,
		)
	}

	fileResponse := this.currentFileResponses[rowIndex]

	/*
		isFolder := fileResponse.IsFolder()
		if isFolder {
			logger.Debugf("FilesPanel.handleFolderClick() - isFolder is true")
		} else {
			logger.Debugf("FilesPanel.handleFolderClick() - isFolder is false")
		}
	*/

	logger.Debugf("FilesPanel.handleFolderClick() - fileResponse name: %s", fileResponse.Name)

	this.locationHistory.GoForward(fileResponse.Name)
	this.doLoadFiles()

	logger.TraceLeave("FilesPanel.handleFolderClick()")
}

func (this *filesPanel) handleFileClick(button *gtk.Button, rowIndex int) {
	logger.TraceEnter("FilesPanel.handleFileClick()")

	logger.Debugf("FilesPanel.handleFileClick() - rowIndex: %d", rowIndex)

	if this.currentFileResponses == nil {
		logger.Fatalf("FilesPanel.handleFileClick() - currentFileResponses is nil")
	}

	fileResponsesLength := len(this.currentFileResponses)
	if rowIndex >= fileResponsesLength {
		logger.Fatalf(
			"FilesPanel.handleFileClick() - rowIndex is out of range.  The length of currentFileResponses is %d but rowIndex is %d",
			fileResponsesLength,
			rowIndex,
		)
	}

	fileResponse := this.currentFileResponses[rowIndex]

	message := ""
	strLen := len(fileResponse.Name)
	if strLen <= 20 {
		message = fmt.Sprintf("Do you wish to print %s?", fileResponse.Name)
	} else {
		message = fmt.Sprintf("Do you wish to print\n%s?", fileResponse.Name)
	}

	uiUtils.MustConfirmDialogBox(this.UI.window, message, func() {
		selectFileRequest := &prusaLinkApis.SelectFileRequest{}

		// Set the location to "local" or "sdcard"
		selectFileRequest.Location = this.locationHistory.CurrentLocation()

		selectFileRequest.Path = fileResponse.Name
		selectFileRequest.Print = true

		logger.Infof("Loading file %q", fileResponse.Name)
		if err := selectFileRequest.Do(this.UI.Client); err != nil {
			logger.LogError("FilesPanel.handleFileClick()", "Do(SelectFileRequest)", err)
			errorMessage := fmt.Sprintf("Failed to print the file:\n\n%s", err.Error())
			uiUtils.ErrorMessageDialogBox(this.UI.window, errorMessage)
			return
		}

		this.UI.GoToPreviousPanel()
	})

	logger.TraceLeave("FilesPanel.handleFileClick()")
}
