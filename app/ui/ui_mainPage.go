package ui_terminal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var mainPageTextString = "Sodium Playlist Manager\n" +
	"(a) to add song\n" +
	"(q) to quit"

var mainPageText = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText(mainPageTextString)

var mainPageFlex = tview.NewFlex()

func mainPageFlexInit() {
	mainPageFlex.Clear()
	mainPageFlex.AddItem(mainPageText, 0, 1, false)
	var listGrid = tview.NewFrame(songList).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText("Playlist Master", true, tview.AlignLeft, tcell.ColorWhite)
	mainPageFlex.AddItem(listGrid, 0, 1, false)

	mainPageFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
		} else if event.Rune() == 'a' {
			runAddSongForm()
			pages.SwitchToPage("addSong")
		}
		return event
	})
}
