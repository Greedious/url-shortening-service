package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Generic function to validate any struct
func ValidateStruct[T any](data T) error {
	err := validate.Struct(data)
	if err != nil {
		return err
	}
	return nil
}
