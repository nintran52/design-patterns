package main

import "fmt"

// Step 1: Define Component Interface
type Coffee interface {
	Cost() float64
	Description() string
}

// Step 2: Define Concrete Component
type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() float64 {
	return 2.0
}

func (c *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// Step 3: Define Base Decorator
type CoffeeDecorator struct {
	coffee Coffee
}

func (d *CoffeeDecorator) Cost() float64 {
	return d.coffee.Cost()
}

func (d *CoffeeDecorator) Description() string {
	return d.coffee.Description()
}

// Step 4: Define Concrete Decorators
type MilkDecorator struct {
	CoffeeDecorator
}

func (m *MilkDecorator) Cost() float64 {
	return m.coffee.Cost() + 0.5
}

func (m *MilkDecorator) Description() string {
	return m.coffee.Description() + ", Milk"
}

type SugarDecorator struct {
	CoffeeDecorator
}

func (s *SugarDecorator) Cost() float64 {
	return s.coffee.Cost() + 0.3
}

func (s *SugarDecorator) Description() string {
	return s.coffee.Description() + ", Sugar"
}

// Step 5: Usage Example
func main() {
	var myCoffee Coffee = &SimpleCoffee{}
	fmt.Println(myCoffee.Description(), "Cost:", myCoffee.Cost())

	myCoffee = &MilkDecorator{CoffeeDecorator{coffee: myCoffee}}
	fmt.Println(myCoffee.Description(), "Cost:", myCoffee.Cost())

	myCoffee = &SugarDecorator{CoffeeDecorator{coffee: myCoffee}}
	fmt.Println(myCoffee.Description(), "Cost:", myCoffee.Cost())
}
