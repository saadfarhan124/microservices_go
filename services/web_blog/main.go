package main

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/saadfarhan124/microservices_go/web_blog.git/database"
	"github.com/saadfarhan124/microservices_go/web_blog.git/models"
)

func main() {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello Universe")
	})
	orm.RegisterModel(new(models.Authors))
	database.ConnectDB()
	log.Fatal(app.Listen(":3069"))
}
