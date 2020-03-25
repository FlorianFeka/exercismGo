package bob

import (
	"strings"
	"regexp"
) 

// Hey ...
func Hey(remark string) string {
	upperRemark := strings.ToUpper(remark)
	isWordRegex, _ := regexp.Compile("[A-Za-z]+")
	isWord := isWordRegex.MatchString(remark)
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	if isWord && remark == upperRemark {
		if string(remark[len(remark)-1]) == "?" {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}

	if string(remark[len(remark)-1]) == "?"{
		return "Sure."
	}
	
	return "Whatever."
}
