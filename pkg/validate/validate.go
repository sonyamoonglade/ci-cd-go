package validate

import "strings"

func ValidateEmail(v string) bool {

	c := 0

	if strings.Contains(v, "@") {
		c += 1
	}
	if strings.Contains(v, ".") {
		c++
	}

	return c == 2
}
