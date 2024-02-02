package ui_terminal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()
var pages = tview.NewPages()
var songs = []string{}
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(a) to add song\n(q) to quit")
var form = tview.NewForm()

var songList = tview.NewList()
var flex = tview.NewFlex()

func flexUpdate() {
	flex.Clear()
	flex.AddItem(text, 0, 1, false)
	var listGrid = tview.NewFrame(songList).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText("Playlist Master", true, tview.AlignLeft, tcell.ColorWhite)
	flex.AddItem(listGrid, 0, 1, false)
}

func addSongForm() {
	var path string
	// form.AddTextView("")
	form.AddInputField("Song Path", "", 20, nil, func(input string) {
		path = input
	})
	form.AddButton("Save", func() {
		// print(path)
		songs = append(songs, path)
		songList.AddItem(path, "idk", rune(48+len(songs)), nil)
		// textUpdate()
		pages.SwitchToPage("Manual")
	})
}

func Run() {

	// страницы
	pages.AddPage("Manual", flex, true, true)
	pages.AddPage("Add Song", form, true, false)

	flexUpdate()

	// обработка команд
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
		} else if event.Rune() == 'a' {
			form.Clear(true)
			addSongForm()
			pages.SwitchToPage("Add Song")
		}
		return event
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
