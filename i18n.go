package i18next

import (
	"errors"
)

type Backend struct {
	LoadPath    []string // from network or local
	CrossDomain bool
}

type I18nOptions struct {
	Lng          []string                     // example: ["en-US", "fr-FR"]
	FallbackLng  []string                     // example: ["en-US"]
	DefaultLng   string                       // example: "en-US"
	Ns           []string                     // example: ["dev", "stage", "prod"], not require. usage: https://example.com/locales/{{.Ns}}/{{.Lng]}.json
	DefaultNS    string                       // example: "dev"
	FallbackNS   []string                     // example: ["stage", "prod"]
	Preload      bool                         // preload lng resource.
	LowerCaseLng bool                         // reverse lng to lowerCase, example: en-US => en-us
	CleanCode    bool                         // reverse lng to clean, example: en-US / en-us => en
	Resources    map[string]map[string]string // load lngs from local, priority: 1
	Backend      Backend                      // load lngs from file or network, priority: 2
}

type I18n struct {
	langs map[string]string
}

func Init(options I18nOptions) I18n {
	i18n := &I18n{}
	if options.Resources != nil {
		i18n.langs = options.Resources[options.DefaultLng]
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
