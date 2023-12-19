package utils

import "strings"

func ToCleanCode(code string) string {
	ca := strings.Split(code, "-")
	return ca[0]
}

func ToLowerCase(code string) string {

	return strings.ToLower(code)
}
