package i18next

import "strings"

type Localizer struct {
}

func (l *Localizer) StringContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (l *Localizer) ToCleanCode(code string) string {
	ca := strings.Split(code, "-")
	return ca[0]
}

func (l *Localizer) ToLowerCase(code string) string {
	return strings.ToLower(code)
}
