package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello Universe")
	})
	log.Fatal(app.Listen(":3069"))
}
