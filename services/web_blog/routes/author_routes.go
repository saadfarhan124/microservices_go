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
	// Get ID from Params and convert it into int
	// The second variable is the err value which we have ignored for now
	id, _ := strconv.Atoi(c.Params("id"))

	// We fetch our base response object
	response := utilities.GetBaseResponseObject()
	// Create an author object using the id we get from url params
	author := &models.Authors{Id: id}
	// Read the author data and check for errors
	if err := database.Database.Orm.Read(author); err != nil {
		response["error"] = err.Error()
		// Check if the error is of type ErrNoRows
		// which basically means that no record was found
		if errors.Is(err, orm.ErrNoRows) {
			response["message"] = "Author not found"
			return c.Status(fiber.StatusNotFound).JSON(response)
		}
		// If the error is of any other type then return 500 code
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	} else {
		// Return author's data if found
		response["data"] = author
		response["status"] = "pass"
		delete(response, "message")
		return c.Status(fiber.StatusOK).JSON(response)
	}
}

func DeleteAuthor(c *fiber.Ctx) error {
	// Get base response object
	response := utilities.GetBaseResponseObject()
	// Initilize postbody variable for delete author
	postBody := &validators.AuthorDeletePostBody{}
	// Use our post body validation function to validate postbody
	// Return 400 if invalid json or postbody
	if err := utilities.PostBodyValidation(c, postBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	} else {
		// Create author object using ID
		author := &models.Authors{Id: postBody.Id}
		// Use delete method of orm layer that either returns
		// The number of rows affected
		// or an error
		// We can use the number of rows var to cater
		// record not found and send the apt code back
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
	// Get base response object
	response := utilities.GetBaseResponseObject()
	// Initilize postbody variable for update author
	postBody := &validators.AuthorUpdatePostBody{}
	// Use our post body validation function to validate postbody
	// Return 400 if invalid json or postbody
	if err := utilities.PostBodyValidation(c, postBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		// Create author object using ID and title from postbody
		author := &models.Authors{Id: postBody.Id, Title: postBody.Title}
		// Use update method of orm layer that either returns
		// The number of rows affected
		// or an error
		// We can use the number of rows var to cater
		// record not found and send the apt code back
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
