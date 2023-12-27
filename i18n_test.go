package i18next_test

import (
	"fmt"
	"github.com/yuangwei/go-i18next"
	"testing"
)

func TestI18nNext(t *testing.T) {
	i18n, err := i18next.Init(i18next.I18nOptions{
		Lng:        []string{"en", "cn"},
		DefaultLng: "en",
		Backend: i18next.Backend{
			LoadPath: []string{
				"./examples/datas/json/{{.Lng}}/home.json",
			},
		},
	})
	if err != nil {
		t.Fatalf("i18n Initial error %s", err)
	}
	val, err := i18n.T("title", struct {
		Name string
	}{Name: "Mike"})

	if err != nil {
		t.Fatalf("i18n Format error %s", err)
	}
	fmt.Println(val)
}
