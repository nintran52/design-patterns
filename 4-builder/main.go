package main

import "fmt"

// Step 1: Define the Product
type Burger struct {
	Bread      string
	Meat       string
	Cheese     string
	Vegetables []string
}

func (b *Burger) Show() {
	fmt.Printf("Burger with %s bread, %s meat, %s cheese, and veggies: %v\n",
		b.Bread, b.Meat, b.Cheese, b.Vegetables)
}

// Step 2: Define the Builder Interface and Concrete Builder
type Builder interface {
	SetBread(bread string)
	SetMeat(meat string)
	SetCheese(cheese string)
	SetVegetables(veggies []string)
	Build() *Burger
}

type BurgerBuilder struct {
	burger *Burger
}

func NewBurgerBuilder() *BurgerBuilder {
	return &BurgerBuilder{burger: &Burger{}}
}

func (b *BurgerBuilder) SetBread(bread string) {
	b.burger.Bread = bread
}

func (b *BurgerBuilder) SetMeat(meat string) {
	b.burger.Meat = meat
}

func (b *BurgerBuilder) SetCheese(cheese string) {
	b.burger.Cheese = cheese
}

func (b *BurgerBuilder) SetVegetables(veggies []string) {
	b.burger.Vegetables = veggies
}

func (b *BurgerBuilder) Build() *Burger {
	return b.burger
}

// Step 3: Define the Director
type Director struct{}

func (d *Director) MakeBeefBurger(builder Builder) {
	builder.SetBread("Sesame")
	builder.SetMeat("Beef")
	builder.SetCheese("Cheddar")
	builder.SetVegetables([]string{"Lettuce", "Tomato"})
}

func (d *Director) MakeVeggieBurger(builder Builder) {
	builder.SetBread("Whole Wheat")
	builder.SetMeat("None")
	builder.SetCheese("Swiss")
	builder.SetVegetables([]string{"Lettuce", "Cucumber", "Onion"})
}

func (d *Director) MakeChickenBurger(builder Builder) {
	builder.SetBread("Brioche")
	builder.SetMeat("Chicken")
	builder.SetCheese("Provolone")
	builder.SetVegetables([]string{"Lettuce", "Tomato", "Pickles"})
}

// Step 4: Client Code
func main() {
	builder := NewBurgerBuilder()
	director := Director{}

	fmt.Println("Building a Beef Burger:")
	director.MakeBeefBurger(builder)
	burger1 := builder.Build()
	burger1.Show()

	fmt.Println("\nBuilding a Veggie Burger:")
	builder = NewBurgerBuilder()
	director.MakeVeggieBurger(builder)
	burger2 := builder.Build()
	burger2.Show()

	fmt.Println("\nBuilding a Chicken Burger:")
	builder = NewBurgerBuilder()
	director.MakeChickenBurger(builder)
	burger3 := builder.Build()
	burger3.Show()
}
