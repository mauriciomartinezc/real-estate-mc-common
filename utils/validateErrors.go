package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"strings"
)

// FormatValidationErrors returns a map where each field has an array of error messages
func FormatValidationErrors(validationErrors validator.ValidationErrors, localice *i18n.Localizer) map[string][]string {
	errors := make(map[string][]string)

	for _, err := range validationErrors {
		field := err.StructField()
		message := fmt.Sprintf("%s %s, %s", transString(localice, locales.InvalidValueFor), transString(localice, field), err.Tag())
		errors[field] = append(errors[parseFieldName(field)], message)
	}

	return errors
}

func transString(localice *i18n.Localizer, string string) string {
	return localice.MustLocalize(&i18n.LocalizeConfig{MessageID: string})
}

func parseFieldName(fieldName string) string {
	if len(fieldName) > 0 {
		fieldName = strings.ToLower(string(fieldName[0])) + fieldName[1:]
	}
	return fieldName
}
