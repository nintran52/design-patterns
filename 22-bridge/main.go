package main

import "fmt"

// Step 1: Define Device interface
type Device interface {
	IsEnabled() bool
	Enable()
	Disable()
	GetVolume() int
	SetVolume(percent int)
}

// Step 2: Define Concrete Devices
type TV struct {
	on     bool
	volume int
}

func (t *TV) IsEnabled() bool { return t.on }
func (t *TV) Enable()         { t.on = true; fmt.Println("TV turned ON") }
func (t *TV) Disable()        { t.on = false; fmt.Println("TV turned OFF") }
func (t *TV) GetVolume() int  { return t.volume }
func (t *TV) SetVolume(v int) {
	t.volume = v
	fmt.Println("TV volume set to", v)
}

type Radio struct {
	on     bool
	volume int
}

func (r *Radio) IsEnabled() bool { return r.on }
func (r *Radio) Enable()         { r.on = true; fmt.Println("Radio turned ON") }
func (r *Radio) Disable()        { r.on = false; fmt.Println("Radio turned OFF") }
func (r *Radio) GetVolume() int  { return r.volume }
func (r *Radio) SetVolume(v int) {
	r.volume = v
	fmt.Println("Radio volume set to", v)
}

// Step 3: Define Remote Control
type Remote struct {
	device Device
}

func NewRemote(d Device) *Remote {
	return &Remote{device: d}
}

func (r *Remote) TogglePower() {
	if r.device.IsEnabled() {
		r.device.Disable()
	} else {
		r.device.Enable()
	}
}

func (r *Remote) VolumeUp() {
	r.device.SetVolume(r.device.GetVolume() + 10)
}

func (r *Remote) VolumeDown() {
	r.device.SetVolume(r.device.GetVolume() - 10)
}

type AdvancedRemote struct {
	*Remote
}

func NewAdvancedRemote(d Device) *AdvancedRemote {
	return &AdvancedRemote{NewRemote(d)}
}

func (r *AdvancedRemote) Mute() {
	fmt.Println("Muting device...")
	r.device.SetVolume(0)
}

// Step 4: Client code
func main() {
	tv := &TV{}
	radio := &Radio{}

	fmt.Println("=== Using TV Remote ===")
	tvRemote := NewAdvancedRemote(tv)
	tvRemote.TogglePower()
	tvRemote.VolumeUp()
	tvRemote.Mute()

	fmt.Println("\n=== Using Radio Remote ===")
	radioRemote := NewRemote(radio)
	radioRemote.TogglePower()
	radioRemote.VolumeUp()
	radioRemote.VolumeDown()
}
