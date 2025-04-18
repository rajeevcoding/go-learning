package media

import (
	"fmt"
)

type Playlist struct {
	Songs []Song
}

func (p *Playlist) AddSong(s Song) {
	p.Songs = append(p.Songs, s)
}

func (p Playlist) TotalDuration() int {
	total := 0
	for _, song := range p.Songs {
		total += song.DurationInSeconds
	}

	return total
}

func (p Playlist) PrintSongs() {
	for _, song := range p.Songs {
		fmt.Printf("Title: %s by %s\n", song.Title, song.Artist)
	}
}
