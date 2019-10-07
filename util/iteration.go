package util

import "strings"

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

func Sum(arr []int) int {
	var n int
	for _, v := range arr {
		n += v
	}
	return n
}
