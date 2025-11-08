package main

import "fmt"

// Step 1 Define Element
type Song struct {
	Title  string
	Artist string
}

// Step 2 Define Iterator Interface
type Iterator interface {
	HasNext() bool
	Next() *Song
}

// Step 3. Define Aggregate (Collection)
type Playlist struct {
	songs []*Song
}

func (p *Playlist) AddSong(song *Song) {
	p.songs = append(p.songs, song)
}

func (p *Playlist) CreateIterator() Iterator {
	return &PlaylistIterator{
		songs: p.songs,
		index: 0,
	}
}

// Step 4. Implement Concrete Iterator
type PlaylistIterator struct {
	songs []*Song
	index int
}

func (it *PlaylistIterator) HasNext() bool {
	return it.index < len(it.songs)
}

func (it *PlaylistIterator) Next() *Song {
	if !it.HasNext() {
		return nil
	}
	song := it.songs[it.index]
	it.index++
	return song
}

func main() {
	playlist := &Playlist{}
	playlist.AddSong(&Song{Title: "Let It Be", Artist: "The Beatles"})
	playlist.AddSong(&Song{Title: "Bohemian Rhapsody", Artist: "Queen"})
	playlist.AddSong(&Song{Title: "Hotel California", Artist: "Eagles"})

	iterator := playlist.CreateIterator()
	fmt.Println("🎵 Playlist:")
	for iterator.HasNext() {
		song := iterator.Next()
		fmt.Printf("- %s by %s\n", song.Title, song.Artist)
	}
}
