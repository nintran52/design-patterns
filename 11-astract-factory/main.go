package main

import "fmt"

// Step 1: Define abstract products
type Button interface {
	Paint()
}

type Checkbox interface {
	Paint()
}

type CheckList interface {
	Paint()
}

// Step 2: Define concrete products
type WindowsButton struct{}

func (b *WindowsButton) Paint() {
	fmt.Println("Rendering a button in Windows style")
}

type MacButton struct{}

func (b *MacButton) Paint() {
	fmt.Println("Rendering a button in Mac style")
}

type WindowsCheckbox struct{}

func (c *WindowsCheckbox) Paint() {
	fmt.Println("Rendering a checkbox in Windows style")
}

type MacCheckbox struct{}

func (c *MacCheckbox) Paint() {
	fmt.Println("Rendering a checkbox in Mac style")
}

type WindowsCheckList struct{}

func (cl *WindowsCheckList) Paint() {
	fmt.Println("Rendering a checklist in Windows style")
}

type MacCheckList struct{}

func (cl *MacCheckList) Paint() {
	fmt.Println("Rendering a checklist in Mac style")
}

// Step 3: Define abstract factory
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
	CreateCheckList() CheckList
}

// Step 4: Define concrete factories
type WindowsFactory struct{}

func (f *WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (f *WindowsFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

func (f *WindowsFactory) CreateCheckList() CheckList {
	return &WindowsCheckList{}
}

type MacFactory struct{}

func (f *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (f *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

func (f *MacFactory) CreateCheckList() CheckList {
	return &MacCheckList{}
}

// Step 5: Client code
type Application struct {
	factory GUIFactory
}

func NewApplication(factory GUIFactory) *Application {
	return &Application{factory: factory}
}

func (app *Application) CreateUI() {
	button := app.factory.CreateButton()
	checkbox := app.factory.CreateCheckbox()
	checklist := app.factory.CreateCheckList()
	button.Paint()
	checkbox.Paint()
	checklist.Paint()
}

// Example usage
func main() {
	fmt.Println("Creating Windows UI")
	winFactory := &WindowsFactory{}
	app := NewApplication(winFactory)
	app.CreateUI()

	fmt.Println("\nCreating Mac UI")
	macFactory := &MacFactory{}
	app = NewApplication(macFactory)
	app.CreateUI()
}
