package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100"`
	Email string `gorm:"uniqueIndex;size:100"`
}

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate
	db.AutoMigrate(&User{})

	// Create
	user := User{Name: "John Doe", Email: "john@example.com"}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// Read
	var foundUser User
	db.First(&foundUser, user.ID)
	fmt.Printf("Found user: %+v\n", foundUser)

	// Update
	db.Model(&foundUser).Update("Name", "Jane Doe")

	// Delete
	db.Delete(&foundUser)

	fmt.Println("GORM operations completed!")
}