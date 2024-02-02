package cmd_handler

import (
	"spm/app/db_handler"
	track_entry "spm/shared/track"
)

var currentPlaylist []track_entry.Track

func AddTrack(track track_entry.Track) bool {
	addToDbSuccess := db_handler.AddTrack(track)
	if addToDbSuccess {
		currentPlaylist = append(currentPlaylist, track)
	}
	return addToDbSuccess
}
