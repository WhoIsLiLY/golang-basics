package main

import "fmt"

func main() {
	// Numbers
	var count int = 42
	var price float64 = 99.99
	fmt.Printf("Count: %d, Price: %.2f\n", count, price)

	// Strings
	message := "Go is awesome"
	fmt.Printf("Message: %s\n", message)

	// Booleans
	isReady := true
	fmt.Printf("Ready: %t\n", isReady)

	// Arrays
	numbers := [3]int{1, 2, 3}
	fmt.Printf("Numbers: %v\n", numbers)

	// Slices
	fruits := []string{"apple", "banana", "orange"}
	fmt.Printf("Fruits: %v\n", fruits)

	// Maps
	scores := map[string]int{
		"Alice": 95,
		"Bob":   87,
	}
	fmt.Printf("Scores: %v\n", scores)
}