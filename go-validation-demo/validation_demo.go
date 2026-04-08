package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserInput struct {
	Name  string `validate:"required,min=2,max=50"`
	Email string `validate:"required,email"`
	Age   int    `validate:"required,min=18,max=100"`
}

func main() {
	validate := validator.New()

	// Valid input
	user1 := UserInput{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   25,
	}

	err := validate.Struct(user1)
	if err != nil {
		fmt.Printf("Validation failed: %v\n", err)
	} else {
		fmt.Println("User1 validation passed!")
	}

	// Invalid input
	user2 := UserInput{
		Name:  "A",
		Email: "invalid-email",
		Age:   15,
	}

	err = validate.Struct(user2)
	if err != nil {
		fmt.Printf("User2 validation failed: %v\n", err)
	} else {
		fmt.Println("User2 validation passed!")
	}
}