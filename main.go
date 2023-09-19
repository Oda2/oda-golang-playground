package main

import (
	"strings"
)

func MyIndexAny(s, chars string) int {
	for i, c := range s {
		if strings.Contains(chars, string(c)) {
			return i
		}
	}

	return -1
}
