package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	abbreviation := ""
	r, _ := regexp.Compile("[A-Za-z\\']+")
	matches := r.FindAllString(s, -1)

	for i := 0; i < len(matches); i++ {
		word := []rune(matches[i])
		abbreviation = abbreviation + strings.ToUpper(string(word[0]))
	}

	return abbreviation
}
