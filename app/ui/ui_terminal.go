package ui_terminal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

var uiFlex *tview.Flex
var uiBody = tview.NewFlex()

func uiFlexInit() {

	var uiHeaderText = tview.NewTextView()
	uiHeaderText.SetText("Sodium Playlist Manager").
		SetTextStyle(tcell.StyleDefault.Bold(true).Background(tcell.ColorDefault)).
		SetTextColor(tcell.ColorRed).
		SetBackgroundColor(tcell.ColorBlack)

	var uiCloseButton = tview.NewButton("Exit")
	uiCloseButton.SetSelectedFunc(func() {
		app.Stop()
	})

	var uiHeader = tview.NewFlex().
		AddItem(uiHeaderText, 0, 1, false).
		AddItem(uiCloseButton, 6, 0, false)

	uiHeader.SetBorderPadding(0, 0, 1, 1)
	uiHeader.SetBackgroundColor(tcell.ColorBlack)
	uiHeader.SetBorder(true)

	uiBody.
		AddItem(getManualFrame(), 0, 1, false).
		AddItem(getBrowserFrame(), 0, 1, false)

	var uiFooter = tview.NewFrame(tview.NewTextView().SetText("Placeholder")).SetBorders(0, 0, 0, 0, 0, 0)
	uiFooter.SetBorder(true)

	uiFlex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(uiHeader, 3, 0, false).
		AddItem(uiBody, 0, 1, false).
		AddItem(uiFooter, 5, 0, false)

}

func OpenPlaylist(name string) bool {
	new_playlist_frame, success := getPlaylistFrame(name)
	if !success {
		return false
	}

	uiBody.AddItem(new_playlist_frame, 0, 1, false)
	return true
	// AddItem(getPlaylistFrame("master"), 0, 1, false).
}

var songList = tview.NewList()

func Run() {
	// mainPageFlexInit()

	uiFlexInit()

	if err := app.SetRoot(uiFlex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
