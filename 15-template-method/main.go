package main

import "fmt"

// Step 1: Define the Template
type Beverage interface {
	PrepareRecipe()
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
}

// Step 2: Create the Base Struct
type BaseBeverage struct{}

func (b *BaseBeverage) BoilWater() {
	fmt.Println("Boiling water")
}

func (b *BaseBeverage) PourInCup() {
	fmt.Println("Pouring into cup")
}

// Step 3: Implement the concrete beverages
type Tea struct {
	BaseBeverage
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding lemon")
}

func (t *Tea) PrepareRecipe() {
	t.BoilWater()
	t.Brew()
	t.PourInCup()
	t.AddCondiments()
}

type Coffee struct {
	BaseBeverage
}

func (c *Coffee) Brew() {
	fmt.Println("Dripping coffee through filter")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding sugar and milk")
}

func (c *Coffee) PrepareRecipe() {
	c.BoilWater()
	c.Brew()
	c.PourInCup()
	c.AddCondiments()
}

// Step 4: Use the Template Method
func main() {
	tea := &Tea{}
	coffee := &Coffee{}

	fmt.Println("Preparing Tea:")
	tea.PrepareRecipe()

	fmt.Println("\nPreparing Coffee:")
	coffee.PrepareRecipe()
}
