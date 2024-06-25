package i18next

import (
	"reflect"
	"testing"
)

func TestBackend_Load_Local(t *testing.T) {
	backend := &Backend{
		LoadPath:    []string{"./examples/json/{{.Lng}}/{{.Ns}}.json"},
		CrossDomain: false,
	}
	langs, err := backend.Load(I18nOptions{
		Lng:         "en-US",
		FallbackLng: "home",
		CleanCode:   true,
	})
	if err != nil {
		t.Errorf("Load from local error, msg: %s", err)
	}
	want := map[string]string{"home": "Hello world"}
	if !reflect.DeepEqual(langs, want) {
		t.Errorf("Backend.Load() got = %v, want %v", langs, want)
	}
}

func TestBackend_Load_Network(t *testing.T) {
	backend := &Backend{
		LoadPath:    []string{"https://example.com/locales/json/{{.Lng}}/{{.Ns}}.json"},
		CrossDomain: false,
	}
	langs, err := backend.Load(I18nOptions{
		Lng:         "en-US",
		FallbackLng: "home",
		CleanCode:   true,
	})
	if err != nil {
		t.Errorf("Load from local error, msg: %s", err)
	}
	want := map[string]string{"home": "Hello world"}
	if !reflect.DeepEqual(langs, want) {
		t.Errorf("Backend.Load() got = %v, want %v", langs, want)
	}
}
