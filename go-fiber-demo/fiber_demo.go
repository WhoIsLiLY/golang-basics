package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.JSON(fiber.Map{
			"user_id": id,
			"message": "User found",
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		type User struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		var user User
		if err := c.BodyParser(&user); err != nil {
			return err
		}

		return c.Status(201).JSON(user)
	})

	log.Fatal(app.Listen(":3000"))
}