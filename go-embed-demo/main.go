package main

import (
	_ "embed"
	"fmt"
)

//go:embed assets/message.txt
var message string

func main() {
	fmt.Println("Embedded content:")
	fmt.Println(message)
}