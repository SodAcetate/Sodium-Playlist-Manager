package cmd_handler

import (
	"spm/app/db_handler"
	spm_types "spm/shared/spm_types"
)

var currentPlaylist []spm_types.Track

func AddTrack(track spm_types.Track) bool {
	addToDbSuccess := db_handler.AddTrack(track)
	if addToDbSuccess {
		currentPlaylist = append(currentPlaylist, track)
	}
	return addToDbSuccess
}

func GetPlaylist(name string) (spm_types.Playlist, bool) {
	// заглушка
	return spm_types.GetPlaceholderPlaylist(name), true
}

func GetPlaylists() []spm_types.PlaylistPreview {
	return spm_types.GetPlaceholderLibrary()
}
