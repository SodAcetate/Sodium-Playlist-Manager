package spm_types

// import "fmt"

type Track struct {
	Path string
	Name string
}

type Playlist struct {
	Name   string
	Tracks []Track
}

type PlaylistPreview struct {
	Name       string
	TrackCount int
}

func GetPlaceholderTrack(name string) Track {
	var song Track
	song.Name = name
	song.Path = "placeholders/" + name
	return song
}

func GetPlaceholderPlaylist(name string) Playlist {
	var playlist Playlist
	playlist.Name = name
	playlist.Tracks = []Track{
		GetPlaceholderTrack("Led Zeppelin -- Heartbreaker"),
		GetPlaceholderTrack("Nirvana -- Smells Like Teen Spirit"),
		GetPlaceholderTrack("Король и Шут -- Лесник"),
	}
	return playlist
}

func GetPlaceholderPlaylistPreview(name string, count int) PlaylistPreview {
	var playlist PlaylistPreview
	playlist.Name = name
	playlist.TrackCount = count
	return playlist
}

func GetPlaceholderLibrary() []PlaylistPreview {
	var library = []PlaylistPreview{
		GetPlaceholderPlaylistPreview("master", 32),
		GetPlaceholderPlaylistPreview("rock", 13),
		GetPlaceholderPlaylistPreview("russian", 10),
		GetPlaceholderPlaylistPreview("eurobeat", 11),
	}
	return library

}

// func toString(Track) string {
// 	return fmt.Sprintf()
// }
