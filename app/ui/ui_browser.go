package ui_terminal

import (
	"fmt"
	"spm/app/cmd_handler"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func getBrowserFrame() *tview.Frame {
	library := cmd_handler.GetPlaylists()

	var playlistHeader = tview.NewTextView().
		SetTextColor(tcell.ColorYellow).
		SetText("All playlists")

	var browserList = tview.NewList()
	for index, playlist := range library {
		browserList.AddItem(playlist.Name,
			fmt.Sprintf("%d tracks",
				playlist.TrackCount),
			rune(49+index),
			nil)
		browserList.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
			OpenPlaylist(s1)
		})
	}

	var browserFlex = tview.NewFlex().SetDirection(tview.FlexRow)

	browserFlex.AddItem(playlistHeader, 1, 0, false)
	browserFlex.AddItem(browserList, 0, 1, false)

	var browserFrame = tview.NewFrame(browserFlex)
	browserFrame.SetBorder(true)

	return browserFrame
}
