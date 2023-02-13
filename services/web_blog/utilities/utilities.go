package utilities

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saadfarhan124/microservices_go/web_blog.git/validators"
)

func GetBaseResponseObject() map[string]interface{} {
	response := make(map[string]interface{})
	response["status"] = "fail"
	response["message"] = "something went wrong"
	return response
}

func PostBodyValidation(c *fiber.Ctx, postBody interface{}) interface{} {
	// Get fiber context and post body interface as an arguement
	// Return an interface back
	// We start by getting the base response object
	response := GetBaseResponseObject()
	// Validate if it is a valid json
	// if not then add error key and send back
	if err := c.BodyParser(postBody); err != nil {
		response["error"] = err.Error()
		return response
	} else {
		// check if post body follows the validations accurately
		// if not then send the errors back
		if errors := validators.ValidateStruct(postBody); errors != nil {
			return errors
		} else {
			// return nil if postbody is both a valid json
			// and passes all validation
			return nil
		}
	}
}
