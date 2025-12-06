package main

import "fmt"

// Sep 1: Define Prototype interface
type Document interface {
	Clone() Document
	ShowInfo()
}

// Sep 2: Define Concrete Prototypes
type Report struct {
	title   string
	content string
}

func (r *Report) Clone() Document {
	copy := *r
	return &copy
}

func (r *Report) ShowInfo() {
	println("Report Title:", r.title)
	println("Report Content:", r.content)
}

// Step 3: Prototype Registry
type DocumentRegistry struct {
	prototypes map[string]Document
}

func NewDocumentRegistry() *DocumentRegistry {
	return &DocumentRegistry{prototypes: make(map[string]Document)}
}

func (dr *DocumentRegistry) Register(name string, doc Document) {
	dr.prototypes[name] = doc
}

func (dr *DocumentRegistry) Create(name string) Document {
	if prototype, exists := dr.prototypes[name]; exists {
		return prototype.Clone()
	}
	fmt.Println("Prototype not found:", name)
	return nil
}

// Step 4: Client code
func main() {
	registry := NewDocumentRegistry()
	reportTemplate := &Report{title: "Monthly Report", content: "Base structure for monthly reports."}
	registry.Register("report", reportTemplate)

	doc1 := registry.Create("report")
	doc1.(*Report).content = "This is the content of the first report."

	doc2 := registry.Create("report")
	doc2.(*Report).content = "This is the content of the second report."

	reportTemplate.ShowInfo()
	doc1.ShowInfo()
	doc2.ShowInfo()
}
