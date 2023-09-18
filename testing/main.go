package main

import (
	"strings"
)

func MyIndexAny(s, chars string) int {
	// p1, _ := utf8.DecodeRuneInString(s)
	// p2, _ := utf8.DecodeRuneInString(chars)

	for i, c := range s {
		if strings.Contains(chars, string(c)) {
			return i
		}
	}

	return -1
}
