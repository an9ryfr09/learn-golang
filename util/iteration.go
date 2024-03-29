package util

import (
	"strings"
)

//Repeat ...
func Repeat(ch string, c int, toUpper bool) string {
	var s string
	for i := 0; i < c; i++ {
		s += ch
	}

	if toUpper {
		return strings.ToUpper(s)
	}
	return s
}
