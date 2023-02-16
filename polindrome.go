package main

import (
	"strings"
	"testing"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"racecar", true},
		{"hello", false},
		{"AbBa", true},
		{"A man a plan a canal Panama", true},
	}
	for _, c := range cases {
		got := IsPalindrome(c.input)
		if got != c.want {
			t.Errorf("IsPalindrome(%q) == %v, want %v", c.input, got, c.want)
		}
	}
}
