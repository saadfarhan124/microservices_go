package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saadfarhan124/microservices_go/web_blog.git/database"
	"github.com/saadfarhan124/microservices_go/web_blog.git/models"
	"github.com/saadfarhan124/microservices_go/web_blog.git/utilities"
	"github.com/saadfarhan124/microservices_go/web_blog.git/validators"
)

func AddAuthor(c *fiber.Ctx) error {
	response := utilities.GetBaseResponseObject()
	postBody := &validators.AuthorAddPostBody{}

	if err := c.BodyParser(postBody); err != nil {
		response["error"] = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	} else {
		if err := validators.ValidateStruct(postBody); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		} else {
			author := models.Authors{Title: postBody.Title}
			if _, err := database.Database.Orm.Insert(&author); err != nil {
				response["error"] = err.Error()
				return c.Status(fiber.StatusInternalServerError).JSON(response)
			} else {
				response["message"] = "Author successfully added"
				response["status"] = "pass"
				return c.Status(fiber.StatusCreated).JSON(response)
			}
		}
	}
}

func GetAllAuthors(c *fiber.Ctx) error {
	return c.SendString("All Author")
}

func GetSingleAuthor(c *fiber.Ctx) error {
	return c.SendString("Single Author")
}

func DeleteAuthor(c *fiber.Ctx) error {
	return c.SendString("Delete Author")
}

func UpdateAuthor(c *fiber.Ctx) error {
	return c.SendString("Update Author")
}
