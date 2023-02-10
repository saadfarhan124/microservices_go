package routes

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/saadfarhan124/microservices_go/web_blog.git/database"
	"github.com/saadfarhan124/microservices_go/web_blog.git/models"
	"github.com/saadfarhan124/microservices_go/web_blog.git/utilities"
	"github.com/saadfarhan124/microservices_go/web_blog.git/validators"
	"strconv"
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
	response := utilities.GetBaseResponseObject()

	qs := database.Database.Orm.QueryTable(models.Authors{}).OrderBy("-created_at")
	var maps []orm.Params

	if count, err := qs.Values(&maps, "id", "title"); err != nil {
		response["error"] = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	} else {
		response["count"] = count
		response["data"] = maps
		response["status"] = "pass"
		delete(response, "message")
		return c.Status(fiber.StatusOK).JSON(response)
	}
}

func GetSingleAuthor(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	response := utilities.GetBaseResponseObject()
	author := &models.Authors{Id: id}
	if err := database.Database.Orm.Read(author); err != nil {
		response["error"] = err.Error()
		if errors.Is(err, orm.ErrNoRows) {
			response["message"] = "Author not found"
			return c.Status(fiber.StatusNotFound).JSON(response)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	} else {
		response["data"] = author
		response["status"] = "pass"
		delete(response, "message")
		return c.Status(fiber.StatusOK).JSON(response)
	}
}

func DeleteAuthor(c *fiber.Ctx) error {
	response := utilities.GetBaseResponseObject()
	postBody := &validators.AuthorDeletePostBody{}
	if err := utilities.PostBodyValidation(c, postBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		author := &models.Authors{Id: postBody.Id}
		if num, err := database.Database.Orm.Delete(author); err != nil {
			response["error"] = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		} else {
			if num == 0 {
				response["message"] = "No record found"
				return c.Status(fiber.StatusNotFound).JSON(response)
			} else {
				response["message"] = "Deleted successfully"
				response["status"] = "pass"
				return c.Status(fiber.StatusOK).JSON(response)
			}
		}
	}
}

func UpdateAuthor(c *fiber.Ctx) error {
	response := utilities.GetBaseResponseObject()
	postBody := &validators.AuthorUpdatePostBody{}
	if err := utilities.PostBodyValidation(c, postBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		author := &models.Authors{Id: postBody.Id, Title: postBody.Title}
		if num, err := database.Database.Orm.Update(author); err != nil {
			response["error"] = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		} else {
			if num == 0 {
				response["message"] = "No record found"
				return c.Status(fiber.StatusNotFound).JSON(response)
			} else {
				response["message"] = "Updated successfully"
				response["status"] = "pass"
				return c.Status(fiber.StatusOK).JSON(response)
			}
		}
	}
}
