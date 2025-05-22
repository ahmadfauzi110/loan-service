package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			result, err := json.Marshal(formatValidationErrors(validationErrors))
			if err != nil {
				return err
			}

			return errors.New(string(result))
		}
		return err
	}
	return nil
}

func formatValidationErrors(ves validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, fe := range ves {
		field := snakeToCamel(fe.Field())
		tag := fe.Tag()

		var msg string
		switch tag {
		case "required":
			msg = fmt.Sprintf("%s is required", field)
		case "email":
			msg = fmt.Sprintf("%s must be a valid email address", field)
		case "oneof":
			msg = fmt.Sprintf("%s must be one of %s", field, fe.Param())
		case "min":
			msg = fmt.Sprintf("%s must be at least %s characters", field, fe.Param())
		default:
			msg = fmt.Sprintf("%s is not valid", field)
		}

		errors[field] = msg
	}
	return errors
}

func snakeToCamel(str string) string {

	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
