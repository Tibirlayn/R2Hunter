package lib

import "strings"

func RemoveEmailSymbols(email string) string {
	var result strings.Builder

	for _, char := range email {
		if char != '@' && char != '.' {
			result.WriteRune(char)
		}
	}

	return result.String()
}