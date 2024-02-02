package ui_terminal

import (
	"spm/app/cmd_handler"
	track_entry "spm/shared/track"

	"github.com/rivo/tview"
)

var addSongForm = tview.NewForm()

func runAddSongForm() {

	addSongForm.Clear(true)

	var track track_entry.Track

	addSongForm.AddTextView("Adding to playlist", "master", 0, 1, false, false)

	addSongForm.AddInputField("Track Path", "", 20, nil, func(input string) {
		track.Path = input
	})

	addSongForm.AddInputField("Track Name", "", 20, nil, func(input string) {
		track.Name = input
	})

	addSongForm.AddButton("Save", func() {
		success := cmd_handler.AddTrack(track)
		if success {
			songList.AddItem(track.Name, track.Path, rune(49+songList.GetItemCount()), nil)
		}
		pages.SwitchToPage("mainPage")
	})

}
