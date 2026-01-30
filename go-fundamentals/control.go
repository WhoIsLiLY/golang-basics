package main

import "fmt"

func main() {
	// If statements
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// Switch
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Friday":
		fmt.Println("TGIF!")
	default:
		fmt.Println("Regular day")
	}

	// For loops
	for i := 0; i < 3; i++ {
		fmt.Printf("Loop %d\n", i)
	}

	// Range
	items := []string{"Go", "Python", "Java"}
	for index, item := range items {
		fmt.Printf("%d: %s\n", index, item)
	}
}