package ui_terminal

import (
	"spm/app/cmd_handler"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func getPlaylistFrame(name string) (*tview.Frame, bool) {
	tracks, success := cmd_handler.GetPlaylist(name)
	if !success {
		return nil, false
	}

	var playlistHeader = tview.NewTextView().
		SetTextColor(tcell.ColorYellow).
		SetText(tracks.Name)

	var playlistHint = tview.NewTextView().
		SetTextColor(tcell.ColorGreen).
		SetText("(q) to close")

	var trackList = tview.NewList()
	for index, track := range tracks.Tracks {
		trackList.AddItem(track.Name, track.Path, rune(49+index), nil)
	}

	var playlistFlex = tview.NewFlex().SetDirection(tview.FlexRow)

	playlistFlex.AddItem(playlistHeader, 1, 0, false)
	playlistFlex.AddItem(playlistHint, 1, 0, false)
	playlistFlex.AddItem(trackList, 0, 1, false)

	var playlistFrame = tview.NewFrame(playlistFlex)
	playlistFrame.SetBorder(true)
	playlistFrame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			uiBody.RemoveItem(playlistFrame)
		}
		return event
	})

	return playlistFrame, true
}
