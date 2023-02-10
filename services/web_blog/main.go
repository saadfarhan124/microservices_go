package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/saadfarhan124/microservices_go/web_blog.git/database"
	"github.com/saadfarhan124/microservices_go/web_blog.git/models"
	"github.com/saadfarhan124/microservices_go/web_blog.git/routes"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
	orm.RegisterModel(new(models.Authors))
	database.ConnectDB()
	app := fiber.New()
	SetupRoutes(app)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func SetupRoutes(app *fiber.App) {
	app.Post("/author", routes.AddAuthor)
	app.Get("/authors", routes.GetAllAuthors)
	app.Get("/author/:id", routes.GetSingleAuthor)
	app.Delete("/author", routes.DeleteAuthor)
	app.Put("/author", routes.UpdateAuthor)
}
