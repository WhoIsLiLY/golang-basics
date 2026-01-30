package main

import "fmt"

func main() {
	// Basic variable declaration
	var username string
	username = "Alice"
	fmt.Printf("User: %s\n", username)

	// Direct initialization
	var age = 25
	fmt.Printf("Age: %d\n", age)

	// Short declaration
	city := "Jakarta"
	fmt.Printf("City: %s\n", city)

	// Multiple variables
	var (
		email    = "alice@example.com"
		isActive = true
	)
	fmt.Printf("Email: %s, Active: %t\n", email, isActive)

	// Reassignment
	city = "Bandung"
	fmt.Printf("New City: %s\n", city)
}