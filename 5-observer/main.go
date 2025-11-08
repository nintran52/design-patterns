package main

import "fmt"

// Step 1: Define the Observer Interface
type UsersObserver interface {
	View(videoNumber float64)
}

type MrBeatSubject interface {
	Subscribe(o UsersObserver)
	Unsubscribe(o UsersObserver)
	Notify()
}

// Step 2: Define the Concrete Subject
type MrBeatChannel struct {
	users       []UsersObserver
	videoNumber float64
}

func (w *MrBeatChannel) Subscribe(o UsersObserver) {
	w.users = append(w.users, o)
}

func (w *MrBeatChannel) Unsubscribe(o UsersObserver) {
	for i, user := range w.users {
		if user == o {
			w.users = append(w.users[:i], w.users[i+1:]...)
			break
		}
	}
}

func (w *MrBeatChannel) Notify() {
	for _, observer := range w.users {
		observer.View(w.videoNumber)
	}
}

func (w *MrBeatChannel) SetVideoNumber(videoNumber float64) {
	fmt.Printf("\nMrBeatChannel: New video number = %.1f\n", videoNumber)
	w.videoNumber = videoNumber
	w.Notify()
}

// Step 3: Define the Concrete Observers
type Hien struct {
	name string
}

func (h *Hien) View(videoNumber float64) {
	fmt.Printf("HienNguyen is viewing video number: %.1f\n", videoNumber)
}

type Quan struct {
	name string
}

func (q *Quan) View(videoNumber float64) {
	fmt.Printf("Quan is viewing video number: %.1f\n", videoNumber)
}

// Step 4: Client Code
func main() {
	channel := &MrBeatChannel{}

	hien := &Hien{name: "Hien"}
	quan := &Quan{name: "Quan"}

	channel.Subscribe(hien)
	channel.Subscribe(quan)

	channel.SetVideoNumber(1.0)
	channel.SetVideoNumber(2.0)

	channel.Unsubscribe(hien)
	channel.SetVideoNumber(3.0)
}
