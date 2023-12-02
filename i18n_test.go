package i18next_test

import (
	"testing"

	"github.com/yuangwei/go-i18next"
)

func TestTranslateBasic(t *testing.T) {
	i18n := i18next.Init(i18next.InitOptions{
		FallbackLng: "en",
		DefaultNS:   "en",
		Resources:   map[string]map[string]string{
			"en": {
				"index.home": "Hello world",
			},
		},
	})

	_, err := i18n.T("index.home", "1", "2")

	if err != nil {
		t.Fatalf(`i18n.T("xxx") = %q,`, err)
	}

}
