package i18next

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
)

type Backend struct {
	LoadPath    []string // from network or local, usage: https://example.com/locales/{{.Lng]}/locale.json
	CrossDomain bool
}

type CodeOptions struct {
	LowerCaseLng bool // reverse lng to lowerCase, example: en-US => en-us
	CleanCode    bool // reverse lng to clean, example: en-US / en-us => en
}

type I18nOptions struct {
	Lng         []string                     // example: ["en-US", "fr-FR"] usage: /locale/{{.Lng}}/locale.json or https://example.com/locales/{{.Lng]}/locale.json
	DefaultLng  string                       // example: "en-US"
	Ns          string                       // example: "dev", , not require. usage: https://example.com/locales/{{.Ns}}/{{.Lng]}.json
	Backend     Backend                      // load lngs from file or network
	Resources   map[string]map[string]string // load lngs from local
	CodeOptions CodeOptions
}

type I18n struct {
	langs map[string]string
	opts  I18nOptions
}

func Init(options I18nOptions) (I18n, error) {
	i18n := &I18n{}
	i18n.opts = options
	langs, err := i18n.loadLangs(i18n.opts.DefaultLng)
	if err != nil {
		return I18n{}, err
	}
	i18n.langs = langs
	return *i18n, nil
}

func (i *I18n) Exist(key string) bool {
	if _, ok := i.langs[key]; ok {
		return true
	}
	return false
}

func (i *I18n) T(key string, props interface{}) (string, error) {
	if v, ok := i.langs[key]; ok {
		tmpl, err := template.New("path").Parse(v)
		if err != nil {
			return "", err
		}
		err = tmpl.Execute(os.Stdout, props)
		return v, nil

	}
	return "", errors.New("Not found")
}

func (i *I18n) ChangeLanguage(lng string) error {
	langs, err := i.loadLangs(lng)
	if err != nil {
		return err
	}
	i.langs = langs
	return nil
}

func (i *I18n) loadLangs(lng string) (map[string]string, error) {
	var resources map[string]map[string]string
	backend := i.opts.Backend
	codeOpts := i.opts.CodeOptions
	lngs := i.opts.Lng
	if !contains(lngs, lng) {
		return nil, errors.New(fmt.Sprintf("%s is not in Lngs", lng))
	}
	if backend.LoadPath != nil {
		res, err := getLangs(i.opts)
		if err != nil {
			return nil, err
		}
		resources = res
	} else {
		resources = i.opts.Resources
	}

	if resources == nil {
		return nil, errors.New("resource is blank")
	}
	return loadResource(lng, resources, codeOpts)
}

func loadResource(lng string, resources map[string]map[string]string, codeOpt CodeOptions) (map[string]string, error) {
	lang := transferLangCode(lng, codeOpt.LowerCaseLng, codeOpt.CleanCode)
	res, ok := resources[lang]
	if ok {
		return res, nil
	}
	return nil, errors.New("resource not found")
}

func getLangs(opts I18nOptions) (map[string]map[string]string, error) {
	var resources map[string]map[string]string
	var op = struct {
		Ns  string
		Lng string
	}{Ns: opts.Ns, Lng: opts.DefaultLng}
	p := opts.Backend.LoadPath[0]
	tmpl, err := template.New("p").Parse(p)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, op)
	if err != nil {
		return nil, err
	}

	data, err := getData(buf.String())
	if err != nil {
		return nil, err
	}

	resData, err := parseData(data, getFilePrefix(p))
	if err != nil {
		return nil, err
	}
	resources = resData
	return resources, nil
}

func transferLangCode(lang string, lowerCase bool, cleanCode bool) string {
	if lowerCase == true {
		lang = toLowerCase(lang)
	}
	if cleanCode == true {
		lang = toCleanCode(lang)
	}
	return lang
}

func getData(url string) ([]byte, error) {
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

func parseData(data []byte, prefix string) (map[string]map[string]string, error) {
	var res map[string]map[string]string
	err := json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func getFilePrefix(fileName string) string {
	all := path.Base(fileName)
	suffix := path.Ext(fileName)
	prefix := all[0 : len(all)-len(suffix)]
	return prefix
}

func toCleanCode(code string) string {
	ca := strings.Split(code, "-")
	return ca[0]
}

func toLowerCase(code string) string {

	return strings.ToLower(code)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
