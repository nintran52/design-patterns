package main

import (
	"fmt"
	"strings"
)

// Step 1: Define the Component interface
type Component interface {
	Show(indent int)
}

// step 2: Create Leaf struct
type File struct {
	name string
}

func (l *File) Show(indent int) {
	fmt.Println(strings.Repeat("  ", indent) + "📄 " + l.name)
}

// Step 3: Create Composite struct
type Folder struct {
	name     string
	children []Component
}

func (f *Folder) Show(indent int) {
	fmt.Println(strings.Repeat("  ", indent) + "📁 " + f.name)
	for _, child := range f.children {
		child.Show(indent + 1)
	}
}

func (f *Folder) Add(child Component) {
	f.children = append(f.children, child)
}

// Step 4: Client code
func main() {
	root := &Folder{name: "Root"}
	docs := &Folder{name: "Documents"}
	pics := &Folder{name: "Pictures"}

	file1 := &File{name: "resume.pdf"}
	file2 := &File{name: "photo.jpg"}
	file3 := &File{name: "notes.txt"}

	docs.Add(file1)
	docs.Add(file3)
	pics.Add(file2)
	root.Add(docs)
	root.Add(pics)

	fmt.Println("📂 File System Structure:")
	root.Show(0)
}
