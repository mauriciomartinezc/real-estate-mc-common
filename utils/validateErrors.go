package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"strings"
	"unicode"
)

// FormatValidationErrors returns a map where each field has an array of error messages
func FormatValidationErrors(localizer *i18n.Localizer, ves validator.ValidationErrors) map[string][]string {
	out := make(map[string][]string)

	for _, fe := range ves {
		ns := fe.StructNamespace() // e.g. "Company.NameTrade"
		tag := fe.Tag()            // e.g. "required" or "max"
		param := fe.Param()        // p. ej. "50" para max=50

		// 1) traduce el nombre del campo
		fieldName, err := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: ns,
		})
		if err != nil {
			fieldName = fe.Field() // fallback GO field
		}

		// 2) traduce la regla
		validationMsg, err := localizer.Localize(&i18n.LocalizeConfig{
			MessageID:    tag,
			TemplateData: map[string]interface{}{"Param": param},
		})
		if err != nil {
			validationMsg = tag
		}

		// 3) junta: "Nombre comercial es obligatorio"
		msg := fmt.Sprintf("%s %s", fieldName, validationMsg)

		// 4) convierte la key para la salida JSON, p. ej. "data_billing.address"
		key := parseFieldName(ns)

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

// parseFieldName toma StructNamespace, p.ej. "Company.DataBilling.City.Name",
// descarta "Company", lo divide y convierte cada parte a snake,
// y finalmente las junta con puntos:
// => "data_billing.city.name"
func parseFieldName(ns string) string {
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
