package main

import "fmt"

// CustomMap represents a map with generic key and value types.
type CustomMap[K comparable, V any] struct {
	data map[K]V
}

// NewCustomMap creates a new instance of CustomMap.
func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

// Insert adds a key-value pair to the CustomMap.
func (m *CustomMap[K, V]) Insert(k K, v V) {
	m.data[k] = v
}

// Get retrieves a value by key from the CustomMap.
func (m *CustomMap[K, V]) Get(k K) (V, bool) {
	val, ok := m.data[k]
	return val, ok
}

func main() {
	// Create a CustomMap with string keys and int values
	customMap := NewCustomMap[string, int]()
	customMap.Insert("Mahmood", 25)
	customMap.Insert("Essi", 30)

	// Retrieve and print values
	if age, found := customMap.Get("Mahmood"); found {
		fmt.Println("Mahmood's age is:", age)
	}
	if age, found := customMap.Get("Essi"); found {
		fmt.Println("Essi's age is:", age)
	}
}
