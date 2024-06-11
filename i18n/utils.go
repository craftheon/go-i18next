package i18n

import "strings"

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func toCleanCode(code string) string {
	ca := strings.Split(code, "-")
	return ca[0]
}

func toLowerCase(code string) string {

	return strings.ToLower(code)
}
