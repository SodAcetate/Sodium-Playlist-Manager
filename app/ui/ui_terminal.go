package ui_terminal

import (
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

var pages = tview.NewPages().
	AddPage("mainPage", mainPageFlex, true, true).
	AddPage("addSong", addSongForm, true, false)

var songList = tview.NewList()

func Run() {
	mainPageFlexInit()

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
