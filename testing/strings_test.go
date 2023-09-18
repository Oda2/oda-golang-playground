package main

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkMyIndexAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyIndexAny("abcdefghijklmnopqrstuvwxyz", "q")
	}
}

func BenchmarkStringsIndexAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.IndexAny("abcdefghijklmnopqrstuvwxyz", "q")
	}
}

func TestMyIndexAny(t *testing.T) {
	tests := []struct {
		s, chars string
		want     int
	}{
		{"abcdefghijklmnopqrstuvwxyz", "q", 16},
		{"", "", -1},
		{"", "a", -1},
		{"a", "", -1},
		{"a", "a", 0},
		{"aaa", "a", 0},
		{"abc", "b", 1},
		{"abc", "xcz", 2},
		{"abc", "xyz", -1},
		{"a\x20c", "plq\x20", 1},
		{"\x00\x01\x02\x03\x04\x05\x06\x07", "\x01\x03\x05\x07\x09\x0b\x0d\x0f", 1},
		{"ab/.,m×©", "×©", 6},
		{"aRegExp*", ".(|)*+?^$[]", 7},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s-%s", test.s, test.chars), func(t *testing.T) {
			if got := MyIndexAny(test.s, test.chars); got != test.want {
				t.Errorf("MyIndexAny(%q, %q) = %d, want %d", test.s, test.chars, got, test.want)
			}
		})
	}
}

func FuzzMyIndexAny(f *testing.F) {
	cases := []struct {
		s, chars string
	}{
		{"abcdefghijklmnopqrstuvwxyz", "q"},
		{"", ""},
		{"", "a"},
		{"a", ""},
		{"a", "a"},
		{"aaa", "a"},
		{"abc", "b"},
		{"abc", "xcz"},
		{"abc", "xyz"},
		{"a\x20c", "plq\x20"},
		{"\x00\x01\x02\x03\x04\x05\x06\x07", "\x01\x03\x05\x07\x09\x0b\x0d\x0f"},
		{"ab/.,m×©", "×©"},
		{"aRegExp*", ".(|)*+?^$[]"},
	}

	for _, test := range cases {
		f.Add(test.s, test.chars)
	}

	f.Fuzz(func(t *testing.T, s, chars string) {
		if got, want := MyIndexAny(s, chars), strings.IndexAny(s, chars); got != want {
			t.Errorf("MyIndexAny(%q, %q) = %d, want %d", s, chars, got, want)
		}
	})
}
