package web

import (
	"github.com/go-playground/validator/v10"
)

func convertErrors(err error) map[string]string {
	validationErrors := err.(validator.ValidationErrors)

	errors := make(map[string]string)

	for _, e := range validationErrors {
		errors[e.Field()] = translateError(e)
	}

	return errors
}

func translateError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "this field is required"
	}

	return "invalid input"
}
