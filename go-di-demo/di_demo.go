package main

import "fmt"

type Database struct {
	URL string
}

func NewDatabase() *Database {
	return &Database{URL: "localhost:5432"}
}

type UserService struct {
	db *Database
}

func NewUserService(db *Database) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetUser(id int) string {
	return fmt.Sprintf("User %d from %s", id, s.db.URL)
}

func main() {
	db := NewDatabase()
	userService := NewUserService(db)
	
	result := userService.GetUser(123)
	fmt.Println(result)
}