package i18next

import (
	"errors"
	"github.com/yuangwei/go-i18next/internal/load"
)

type Backend struct {
	LoadPath    []string // from network or local
	CrossDomain bool
}

type CodeOptions struct {
	LowerCaseLng bool // reverse lng to lowerCase, example: en-US => en-us
	CleanCode    bool // reverse lng to clean, example: en-US / en-us => en
}

type I18nOptions struct {
	Lng         []string                     // example: ["en-US", "fr-FR"]
	FallbackLng []string                     // example: ["en-US"]
	DefaultLng  string                       // example: "en-US"
	NS          []string                     // example: ["dev", "prod"]
	DefaultNs   string                       // example: "dev", , not require. usage: https://example.com/locales/{{.Ns}}/{{.Lng]}.json
	Preload     bool                         // preload lng resource.
	Backend     Backend                      // load lngs from file or network
	Resources   map[string]map[string]string // load lngs from local
	CodeOptions CodeOptions
}

type I18n struct {
	langs map[string]string
}

func Init(options I18nOptions) I18n {
	i18n := &I18n{}
	if options.Resources != nil {
		langs, err := load.LoadResourceInLocal(
			options.DefaultLng,
			options.FallbackLng,
			options.Resources,
			options.CodeOptions.LowerCaseLng,
			options.CodeOptions.CleanCode,
		)
		if err != nil {
			panic(any(err))
		}
		i18n.langs = langs
	}
	return *i18n
}

func (i *I18n) Exist(key string) bool {
	// tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
}

func (i *I18n) T(key string, props ...string) (string, error) {
	if v, ok := i.langs[key]; ok {
		return v, nil

	}
	return "", errors.New("Not found")
}

func (i *I18n) ChangeLanguage(lang string) error {

}
