package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello Go Programming"
	
	fmt.Println("Original:", text)
	fmt.Println("Upper:", strings.ToUpper(text))
	fmt.Println("Lower:", strings.ToLower(text))
	fmt.Println("Contains 'Go':", strings.Contains(text, "Go"))
	fmt.Println("Split:", strings.Split(text, " "))
	fmt.Println("Replace:", strings.Replace(text, "Go", "Golang", 1))
}