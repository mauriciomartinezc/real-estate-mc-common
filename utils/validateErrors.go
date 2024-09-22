package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// FormatValidationErrors returns a map where each field has an array of error messages
func FormatValidationErrors(validationErrors validator.ValidationErrors) map[string][]string {
	errors := make(map[string][]string)

	for _, err := range validationErrors {
		field := err.StructField()
		message := fmt.Sprintf("Invalid value for %s, failed %s validation", field, err.Tag())
		errors[field] = append(errors[field], message)
	}

	return errors
}
