package helper

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return nil
	}

	errorsMap := make(map[string]string, len(ve))
	for _, fe := range ve {
		errorsMap[fe.Field()] = msgForTag(fe)
	}

	return errorsMap
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email address format"
	case "min":
		return fmt.Sprintf("Must be at least %s characters long", fe.Param())
	case "max":
		return fmt.Sprintf("Must not exceed %s characters", fe.Param())
	default:
		return "Invalid value"
	}
}
