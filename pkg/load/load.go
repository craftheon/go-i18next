package load

import (
	"errors"
	"github.com/yuangwei/go-i18next/pkg/utils"
)

func LoadResourceInLocal(lang string, fallbackLng []string, resources map[string]map[string]string, lowerCase bool, cleanCode bool) (map[string]string, error) {
	lang = transferLangCode(lang, lowerCase, cleanCode)
	res, ok := loadResource(lang, resources)
	if ok {
		return res, nil
	}
	for _, lng := range fallbackLng {
		res, ok := loadResource(lng, resources)
		if ok {
			return res, nil
		}
	}
	return nil, errors.New("corpus not found")
}

func loadResource(lang string, resources map[string]map[string]string) (map[string]string, bool) {
	val, ok := resources[lang]
	return val, ok
}

func transferLangCode(lang string, lowerCase bool, cleanCode bool) string {
	if lowerCase == true {
		lang = utils.ToLowerCase(lang)
	}
	if cleanCode == true {
		lang = utils.ToCleanCode(lang)
	}
	return lang
}
