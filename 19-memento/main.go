package main

import "fmt"

// Step 1: Define Memento
type Memento struct {
	state string
}

func (m *Memento) GetState() string {
	return m.state
}

// Step 2: Define Originator
type TextEditor struct {
	content string
}

func (t *TextEditor) Type(word string) {
	t.content += word
}

func (t *TextEditor) Save() *Memento {
	fmt.Println("Saving state:", t.content)
	return &Memento{state: t.content}
}

func (t *TextEditor) Restore(m *Memento) {
	t.content = m.GetState()
}

func (t *TextEditor) GetContent() string {
	return t.content
}

// Step 3: Define Caretaker
type History struct {
	snapshots []*Memento
}

func (h *History) Save(m *Memento) {
	h.snapshots = append(h.snapshots, m)
}

func (h *History) Undo() *Memento {
	if len(h.snapshots) == 0 {
		return nil
	}
	last := h.snapshots[len(h.snapshots)-1]
	h.snapshots = h.snapshots[:len(h.snapshots)-1]
	return last
}

func main() {
	editor := &TextEditor{}
	history := &History{}

	editor.Type("Hello ")
	history.Save(editor.Save())

	editor.Type("world! ")
	history.Save(editor.Save())

	editor.Type("This won't be saved.")

	fmt.Println("Current content:", editor.GetContent())

	// Undo last change
	editor.Restore(history.Undo())
	fmt.Println("After undo:", editor.GetContent())

	// Undo again
	editor.Restore(history.Undo())
	fmt.Println("After second undo:", editor.GetContent())
}
