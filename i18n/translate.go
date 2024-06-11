package i18n

import "bytes"

func (i *I18n) T(key string, options interface{}) string {
	if v, ok := i.langs[key]; ok {
		res, err := parseKey(v, options)
		if err != nil {
			_ = eventBus.Emit("missingKey")
		}
		return res
	}
	_ = eventBus.Emit("missingKey")
	return ""
}

func (i *I18n) Exists(key string) bool {
	if _, ok := i.langs[key]; ok {
		return true
	}
	return false
}

func (i *I18n) GetFixedT(lng string, ns string, keyPrefix string) string {

}

func parseKey(content string, props interface{}) (string, error) {
	tmpl, err := template.New("path").Parse(content)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, props)
	return buf.String(), nil
}
