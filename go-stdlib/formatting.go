package main

import "fmt"

func main() {
	name := "Emma"
	age := 28

	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Sprintf("Formatted: %s is %d years old", name, age)
}