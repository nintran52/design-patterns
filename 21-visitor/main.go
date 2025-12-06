package main

import "fmt"

// Step 1: Define Interface
// Element
type Element interface {
	Accept(v Visitor)
}

// Visitor
type Visitor interface {
	VisitParagraph(p *Paragraph)
	VisitImage(i *Image)
	VisitTable(t *Table)
}

// Step 2: Define Concrete Elements
type Paragraph struct {
	text string
}

func (p *Paragraph) Accept(v Visitor) {
	v.VisitParagraph(p)
}

type Image struct {
	path string
}

func (i *Image) Accept(v Visitor) {
	v.VisitImage(i)
}

type Table struct {
	rows int
}

func (t *Table) Accept(v Visitor) {
	v.VisitTable(t)
}

// Step 3: Define Concrete Visitor
type WordCountVisitor struct {
	totalWords int
}

func (w *WordCountVisitor) VisitParagraph(p *Paragraph) {
	// Simple word counting logic
	count := len(splitWords(p.text))
	w.totalWords += count
	fmt.Printf("Paragraph: \"%s\" has %d words\n", p.text, count)
}

func (w *WordCountVisitor) VisitImage(i *Image) {
	fmt.Printf("Skipping image: %s\n", i.path)
}

func (w *WordCountVisitor) VisitTable(t *Table) {
	fmt.Printf("Skipping table with %d rows\n", t.rows)
}

func splitWords(s string) []string {
	var words []string
	word := ""
	for _, ch := range s {
		if ch == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(ch)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}

type HTMLExportVisitor struct{}

func (h *HTMLExportVisitor) VisitParagraph(p *Paragraph) {
	fmt.Printf("<p>%s</p>\n", p.text)
}

func (h *HTMLExportVisitor) VisitImage(i *Image) {
	fmt.Printf("<img src='%s' />\n", i.path)
}

func (h *HTMLExportVisitor) VisitTable(t *Table) {
	fmt.Printf("<table><tr><td>Table with %d rows</td></tr></table>\n", t.rows)
}

// Step 4: Client code
func main() {
	doc := []Element{
		&Paragraph{text: "The visitor pattern separates algorithm from data structure."},
		&Image{path: "diagram.png"},
		&Paragraph{text: "It promotes open/closed principle."},
		&Table{rows: 3},
	}

	counter := &WordCountVisitor{}
	exporter := &HTMLExportVisitor{}
	for _, e := range doc {
		e.Accept(counter)
		e.Accept(exporter)
	}

	fmt.Printf("\nTotal word count: %d\n", counter.totalWords)
}
