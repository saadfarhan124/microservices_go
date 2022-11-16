package validators

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	FailedField string
	Tag         string
}

type AuthorAddPostBody struct {
	Title string `json:"title" validate:"required"`
}

var validate = validator.New()

func ValidateStruct(postBody interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(postBody)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
	}
	return errors
}
