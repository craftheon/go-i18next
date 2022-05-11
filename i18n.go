package i18next

import "fmt"

type Backend struct {
	LoadPath []string
}

type InitOptions struct {
	FallbackLng string
	Ns          []string
	DefaultNS   string
	Debug       bool
	Resource    interface{}
	Backend     Backend
}

type I18n struct{}

type Trans func(key string, props interface{}) string

type InitCallback func(err, t Trans) string

func Init(options InitOptions) I18n {
	fmt.Println(options)
	return I18n{}
}

func Exist(key string) {

}

func T(key string, props ...string) string {
	fmt.Println(props)
	return key
}

func ChangeLanguage(lang string) {

}
