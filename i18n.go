package i18next

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
