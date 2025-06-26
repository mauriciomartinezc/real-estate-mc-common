package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"strings"
	"unicode"
)

// FormatValidationErrors returns a map where each field has an array of error messages
func FormatValidationErrors(localice *i18n.Localizer, validationErrors validator.ValidationErrors) map[string][]string {
	out := make(map[string][]string)

	for _, fe := range validationErrors {
		ns := fe.StructNamespace()

		// intentamos traducir el msg completo a través de i18n
		localized, err := localice.Localize(&i18n.LocalizeConfig{
			MessageID: ns,
		})
		var msg string
		if err != nil {
			// fallback genérico
			msg = fmt.Sprintf("%s %s %s",
				localice.MustLocalize(&i18n.LocalizeConfig{MessageID: "InvalidValueFor"}),
				fe.Field(),
				fe.Tag(),
			)
		} else {
			msg = localized
		}

		key := formatKey(ns)
		out[key] = append(out[key], msg)
	}

	return out
}

// toSnake convierte un sólo segmento CamelCase a snake_case.
func toSnake(s string) string {
	var sb strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				sb.WriteRune('_')
			}
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// formatKey toma StructNamespace, p.ej. "Company.DataBilling.City.Name",
// descarta "Company", lo divide y convierte cada parte a snake,
// y finalmente las junta con puntos:
// => "data_billing.city.name"
func formatKey(ns string) string {
	parts := strings.Split(ns, ".")
	if len(parts) <= 1 {
		return toSnake(ns)
	}
	// ignoramos el primer segmento
	segs := parts[1:]
	for i, seg := range segs {
		segs[i] = toSnake(seg)
	}
	return strings.Join(segs, ".")
}
