package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"regexp"
	"strings"
)

// FormatValidationErrors returns a map where each field has an array of error messages
func FormatValidationErrors(localice *i18n.Localizer, validationErrors validator.ValidationErrors) map[string][]string {
	errors := make(map[string][]string)

	for _, err := range validationErrors {
		field := err.StructField()
		message := fmt.Sprintf("%s %s, %s", transString(localice, locales.InvalidValueFor), transString(localice, field), transString(localice, err.Tag()))
		errors[parseFieldName(field)] = append(errors[parseFieldName(field)], message)
	}

	return errors
}

func transString(localice *i18n.Localizer, string string) string {
	return localice.MustLocalize(&i18n.LocalizeConfig{MessageID: string})
}

func parseFieldName(fieldName string) string {
	regex := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := regex.ReplaceAllString(fieldName, "${1}_${2}")

	return strings.ToLower(snake)
}
