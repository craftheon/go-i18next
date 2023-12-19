package load

import "testing"

func TestLoadResourceLocal(t *testing.T) {
	langs := map[string]map[string]string{
		"en": {
			"index.home": "Hello world",
		},
	}
	_, err := LoadResourceInLocal("en-US", []string{"en-US", "zh-CN"}, langs, true, true)
	if err != nil {
		t.Fatalf(`LoadResourceIn = %q,`, err)
	}
}
