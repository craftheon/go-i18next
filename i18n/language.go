package i18n

func (i *I18n) ChangeLanguage(lng string) error {
	err := eventBus.Emit("languageChanged")
	if err != nil {

	}
}

func (i *I18n) LoadNamespaces() error {

}

func (i *I18n) LoadLanguages() error {

}

func (i *I18n) SetDefaultNamespace() error {}

func (i *I18n) Dir(lng string) {

}
