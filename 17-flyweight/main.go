package main

import "fmt"

// Step 1: Define the Flyweight (Shared Object)
type Character struct {
	symbol string
}

func (c *Character) Display(font string, color string, position int) {
	fmt.Printf("Displaying '%s' with font=%s, color=%s at position=%d\n",
		c.symbol, font, color, position)
}

// Step 2: Create the Flyweight Factory (Object Pool)
type CharacterFactory struct {
	characters map[string]*Character
}

func NewCharacterFactory() *CharacterFactory {
	return &CharacterFactory{characters: make(map[string]*Character)}
}

func (f *CharacterFactory) GetCharacter(symbol string) *Character {
	if _, exists := f.characters[symbol]; !exists {
		fmt.Println("Creating new character:", symbol)
		f.characters[symbol] = &Character{symbol: symbol}
	} else {
		fmt.Println("Reusing existing character:", symbol)
	}
	return f.characters[symbol]
}

// Step 3: Client code
func main() {
	factory := NewCharacterFactory()

	text := []string{"A", "B", "A", "C", "A", "B"}

	position := 0
	for _, ch := range text {
		char := factory.GetCharacter(ch)
		char.Display("Arial", "Black", position)
		position++
	}
}
