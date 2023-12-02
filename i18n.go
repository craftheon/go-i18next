package i18next

import (
	"errors"
)

type Backend struct {
	LoadPath []string // from network or local
	CrossDomain bool
}

type InitOptions struct {
	FallbackLng string
	Ns          []string
	DefaultNS   string
	Debug       bool
	Resources   map[string]map[string]string
	Backend     Backend
}

type I18n struct{
	langMap map[string]string
}


func Init(options InitOptions) I18n {
	i18n := &I18n{}
	if options.Resources != nil {
		i18n.langMap = options.Resources[options.DefaultNS]
	}
	return *i18n
}

func (i *I18n) Exist(key string) {

}

func (i *I18n) T(key string, props ...string) (string, error) {
	if v, ok := i.langMap[key]; ok {
		return v, nil
		
	}
	return "", errors.New("Not found")
}

func (i *I18n) ChangeLanguage(lang string) {

}
