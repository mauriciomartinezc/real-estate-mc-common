package i18n

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
)

func NewLocalization() *i18n.Bundle {
	cwd, _ := os.Getwd()
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile(filepath.Join(cwd, "/i18n/locales/en.json"))
	bundle.MustLoadMessageFile(filepath.Join(cwd, "/i18n/locales/es.json"))
	bundle.MustLoadMessageFile(filepath.Join(cwd, "/i18n/locales/validation_en.json"))
	bundle.MustLoadMessageFile(filepath.Join(cwd, "/i18n/locales/validation_es.json"))
	return bundle
}
