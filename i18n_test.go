package i18next_test

import (
	"github.com/yuangwei/go-i18next"
	"testing"
)

func TestTranslate(t *testing.T) {
	i18next.Init(i18next.InitOptions{
		DefaultNS: "en",
		Backend: i18next.Backend{
			LoadPath: []string{
				"https://www.baidu.com",
			},
		},
	})

	f := i18next.T("xxx", "1", "2")

	if f != "xxx" {
		t.Fatalf(`i18n.T("xxx") = %q,`, f)
	}
}
