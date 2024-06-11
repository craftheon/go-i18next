package i18n

var eventBus EventBus

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

func Init(options I18nOptions, uses ...[]interface{}) (*I18n, error) {
	i18n := &I18n{}
	i18n.options = options
	data, err := loadResources(options)
	if err != nil {
		_ = eventBus.Emit("failedLoading", err)
	}
	_ = eventBus.Emit("loaded", data)
	i18n.langs = data
	_ = eventBus.Emit("initialized", options)
	return i18n, nil
}

func On(event string, callback func(options interface{})) error {
	err := eventBus.Add(event, callback)
	if err != nil {
		return err
	}
	return nil
}
