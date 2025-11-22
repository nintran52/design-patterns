package main

import "fmt"

// Step 1: Define Interface
type Image interface {
	Display()
}

// Step 2: Real Subject
type RealImage struct {
	filename string
}

func (r *RealImage) LoadFromDisk() {
	println("Loading image from disk: " + r.filename)
}

func (r *RealImage) Display() {
	println("Displaying image: " + r.filename)
}

// Step 3: Proxy
type ProxyImage struct {
	realImage *RealImage
	filename  string
}

func (p *ProxyImage) Display() {
	if p.realImage == nil {
		fmt.Println("(Proxy) Image not loaded yet. Loading now...")
		p.realImage = &RealImage{filename: p.filename}
		p.realImage.LoadFromDisk()
	} else {
		fmt.Println("(Proxy) Image already loaded. Using cached version.")
	}
	p.realImage.Display()
}

// Client code
func main() {
	image := &ProxyImage{filename: "test_image.jpg"}

	// Image will be loaded from disk
	image.Display()

	fmt.Println()

	// Image will be displayed from cache
	image.Display()
}
