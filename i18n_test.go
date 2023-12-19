package i18next_test

import (
	"fmt"
	"testing"

	"github.com/yuangwei/go-i18next"
)

func TestLngInLocal(t *testing.T) {
	i18n := i18next.Init(i18next.I18nOptions{
		Lng: []string{"en"},
		Resources: map[string]map[string]string{
			"en": {
				"index.home": "Hello world",
			},
		},
	})
	text, err := i18n.T("index.home")
	if err != nil {
		t.Fatalf(`i18n.T("xxx") = %q,`, err)
	}
	fmt.Println(text)
}

func TestLngInFile(t *testing.T) {
	i18n := i18next.Init(i18next.I18nOptions{
		Lng: []string{"en"},
		Backend: i18next.Backend{
			LoadPath: []string{"/locale/{{.Lng}}/index.json"},
		},
	})
	text, err := i18n.T("index.home")
	if err != nil {
		t.Fatalf(`i18n.T("index.home") = %q,`, err)
	}
	fmt.Println(text)
}

func TestLngInHttp(t *testing.T) {
	i18n := i18next.Init(i18next.I18nOptions{
		Lng: []string{"en"},
		Backend: i18next.Backend{
			LoadPath: []string{"https://yuangwei.com/example-data/go-i18next/locale/{{.Lng}}/index.json"},
		},
	})
	text, err := i18n.T("index.home")
	if err != nil {
		t.Fatalf(`i18n.T("index.home") = %q,`, err)
	}
	fmt.Println(text)
}

func TestLngMutiFile(t *testing.T) {
	i18n := i18next.Init(i18next.I18nOptions{
		Lng: []string{"en"},
		Backend: i18next.Backend{
			LoadPath: []string{"/locale/{{.Lng}}/index.json", "/locale/{{.Lng}}/home.json"},
		},
	})
	text, err := i18n.T("index.home")
	if err != nil {
		t.Fatalf(`i18n.T("index.home") = %q,`, err)
	}
	text, err = i18n.T("home.home")
	if err != nil {
		t.Fatalf(`i18n.T("home.home") = %q,`, err)
	}
	fmt.Println(text)
}

func TestChangeLng(t *testing.T) {
	i18n := i18next.Init(i18next.I18nOptions{
		Lng:        []string{"en", "jp"},
		DefaultLng: "en",
		Backend: i18next.Backend{
			LoadPath: []string{"/locale/{{.Lng}}/index.json"},
		},
	})
	_, err := i18n.T("index.home")
	if err != nil {
		t.Fatalf(`i18n.T("index.home") = %q,`, err)
	}
	err = i18n.ChangeLanguage("jp")
	if err != nil {
		t.Fatalf(`i18n.ChangeLanguage("jp") = %q`, err)
	}
	text, e := i18n.T("index.home")
	if e != nil {
		t.Fatalf(`i18n.T("index.home // jp") = %q,`, err)
	}
	fmt.Println(text)
}

func TestNs(t *testing.T) {
	i18n := i18next.Init(i18next.I18nOptions{
		Lng:       []string{"en"},
		DefaultNS: "dev",
		Backend: i18next.Backend{
			LoadPath: []string{"/locale/{{.Ns}}/{{.Lng}}/index.json"},
		},
	})
	_, err := i18n.T("index.home")
	if err != nil {
		t.Fatalf(`i18n.T("index.home") = %q,`, err)
	}
}
