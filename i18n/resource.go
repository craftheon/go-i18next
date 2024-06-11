package i18n

func (i *I18n) ReloadResources() error {
	data, err := loadResources(i.options)
	if err != nil {
		_ = eventBus.Emit("failedLoading", err)
		return err
	}
	_ = eventBus.Emit("loaded", data)
	i.langs = data
	return nil
}

func loadResources(options I18nOptions) (map[string]string, error) {
	resources := options.Resources
	if resources != nil {

	}
}
