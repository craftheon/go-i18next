package i18next

type Backend struct {
	LoadPath    []string // from network or local, usage: https://example.com/locales/{{.Lng]}/locale.json
	CrossDomain bool
}

func (b *Backend) Load(options I18nOptions) (map[string]string, error) {

}

func (b *Backend) Reload() {

}
