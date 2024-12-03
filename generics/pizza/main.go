package main

import "fmt"

// Pizza represents a generic pizza structure.
type Pizza[T any] struct {
	Name       string
	Ingredients []T
	Price      float64
}

// NewPizza creates a new instance of Pizza.
func NewPizza[T any](name string, ingredients []T, price float64) *Pizza[T] {
	return &Pizza[T]{
		Name:        name,
		Ingredients: ingredients,
		Price:       price,
	}
}

// PrintDetails prints the pizza's details.
func (p *Pizza[T]) PrintDetails() {
	fmt.Printf("Pizza: %s\n", p.Name)
	fmt.Println("Ingredients:")
	for _, ingredient := range p.Ingredients {
		fmt.Printf("- %v\n", ingredient)
	}
	fmt.Printf("Price: $%.2f\n", p.Price)
}

func main() {
	// Create a pizza with string ingredients
	margherita := NewPizza("Margherita", []string{"Tomato", "Mozzarella", "Basil"}, 8.99)
	margherita.PrintDetails()

	// Create a pizza with custom ingredient types
	type CustomIngredient struct {
		Name  string
		WeightInGrams int
	}
	pepperoni := NewPizza("Pepperoni", []CustomIngredient{
		{Name: "Pepperoni", WeightInGrams: 50},
		{Name: "Cheese", WeightInGrams: 100},
	}, 12.50)
	pepperoni.PrintDetails()
}
