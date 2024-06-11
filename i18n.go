package i18next

import "errors"

type I18n struct {
	langs     map[string]string
	options   I18nOptions
	Language  string
	Languages []string
}

type I18nOptions struct {
	Lng          string
	FallbackLng  string
	Load         []string
	PreLoad      bool
	LowerCaseLng bool
	CleanCode    bool
	Ns           []string
	DefaultNs    string
	fallbackNS   bool
	Resources    map[string]map[string]string
	Backend      Backend
}

var (
	translater Translater
	events     Event
)

func Init(options I18nOptions, uses ...[]interface{}) (*I18n, error) {
	i18n := &I18n{}
	i18n.options = options
	if options.Resources != nil {
		res, ok := options.Resources[options.Lng]
		if !ok {
			err := errors.New("") // TODO
			_ = events.Emit("failedLoading", err)
			return nil, err
		}
		i18n.langs = res
		_ = events.Emit("loaded", res)
	} else {
		res, err := i18n.options.Backend.Load(i18n.options)
		if err != nil {
			_ = events.Emit("failedLoading", err)
			return nil, err
		}
		i18n.langs = res
		_ = events.Emit("loaded", res)
	}
	_ = events.Emit("initialized", options)
	return i18n, nil
}

func On(event string, callback func(options interface{})) error {
	err := events.Add(event, callback)
	if err != nil {
		return err
	}
	return nil
}

func (i *I18n) T(key string, options interface{}) string {
	return translater.Translate(key, options)
}

func (i *I18n) ReloadResource() {

}

func (i *I18n) GetFixedT() {

}

func (i *I18n) ChangeLanguage() {

}

func (i *I18n) LoadLanguages() {

}

func (i *I18n) Format() {

}
