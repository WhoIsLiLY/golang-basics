package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Encoding to JSON
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error encoding: %v\n", err)
		return
	}
	fmt.Printf("JSON: %s\n", string(jsonData))

	// Decoding from JSON
	jsonString := `{"id":2,"name":"Jane Smith","email":"jane@example.com"}`
	var newUser User
	
	err = json.Unmarshal([]byte(jsonString), &newUser)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
		return
	}
	fmt.Printf("User: %+v\n", newUser)
}