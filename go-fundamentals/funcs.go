package main

import "fmt"

func greet() {
	fmt.Println("Welcome to Go!")
}

func greetUser(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(a, b int) int {
	return a + b
}

func getInfo() (string, int) {
	return "Bob", 30
}

func main() {
	greet()
	greetUser("Charlie")
	
	result := add(10, 20)
	fmt.Printf("Sum: %d\n", result)
	
	name, age := getInfo()
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}