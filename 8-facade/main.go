package main

import "fmt"

// Step 1. Define Subsystem Components
type DVDPlayer struct{}

func (d *DVDPlayer) On()               { fmt.Println("DVD Player On") }
func (d *DVDPlayer) Play(movie string) { fmt.Println("Playing movie:", movie) }
func (d *DVDPlayer) Off()              { fmt.Println("DVD Player Off") }

type Projector struct{}

func (p *Projector) On()  { fmt.Println("Projector On") }
func (p *Projector) Off() { fmt.Println("Projector Off") }

type Amplifier struct{}

func (a *Amplifier) On()                 { fmt.Println("Amplifier On") }
func (a *Amplifier) SetVolume(level int) { fmt.Println("Set volume to", level) }
func (a *Amplifier) Off()                { fmt.Println("Amplifier Off") }

type Light struct{}

func (l *Light) Dim(level int) { fmt.Println("Lights dimmed to", level, "%") }
func (l *Light) On()           { fmt.Println("Lights On") }

type Screen struct{}

func (s *Screen) Down() { fmt.Println("Screen down") }
func (s *Screen) Up()   { fmt.Println("Screen up") }

// Step 2. Define Facade
type HomeTheaterFacade struct {
	dvd    *DVDPlayer
	amp    *Amplifier
	proj   *Projector
	light  *Light
	screen *Screen
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		dvd:    &DVDPlayer{},
		amp:    &Amplifier{},
		proj:   &Projector{},
		light:  &Light{},
		screen: &Screen{},
	}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("🎬 Get ready to watch a movie...")
	h.light.Dim(10)
	h.screen.Down()
	h.proj.On()
	h.amp.On()
	h.amp.SetVolume(8)
	h.dvd.On()
	h.dvd.Play(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("🎞️ Shutting down theater...")
	h.light.On()
	h.screen.Up()
	h.proj.Off()
	h.amp.Off()
	h.dvd.Off()
}

// Step 3. Client Code
func main() {
	homeTheater := NewHomeTheaterFacade()
	homeTheater.WatchMovie("Inception")
	fmt.Println()
	homeTheater.EndMovie()
}
