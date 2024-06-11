package i18n

import (
	"encoding/json"
	"errors"
	"github.com/pelletier/go-toml/v2"
	"io"
	"net/http"
	"os"
	"strings"
)

type Backend struct {
	LoadPath    []string // from network or local, usage: https://example.com/locales/{{.Lng]}/locale.json
	CrossDomain bool
}

func getResourceFromBackend(url string) ([]byte, error) {
	if strings.Contains(url, "http://") || strings.Contains(url, "https://") {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(resp.Body)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
	resp, err := os.ReadFile(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func decodeResourceFile(data []byte, lng string, prefix string) (map[string]map[string]string, error) {
	var res map[string]string
	var err error
	if prefix == ".json" {
		err = json.Unmarshal(data, &res)
	} else if prefix == ".yaml" || prefix == ".yml" {
		err = yaml.Unmarshal(data, &res)
	} else if prefix == ".toml" {
		err = toml.Unmarshal(data, &res)
	} else {
		return nil, errors.New("prefix not support")
	}
	if err != nil {
		return nil, err
	}
	resource := map[string]map[string]string{}
	resource[lng] = res
	return resource, nil
}
