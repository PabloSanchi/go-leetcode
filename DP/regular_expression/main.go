package main

import (
	"fmt"
)

func dp(s, p string, i, j int) bool {
	if i >= len(s) && j >= len(p) {
		return true
	}
	if j >= len(p) {
		return false
	}

	match := i < len(s) && (p[j] == '.' || s[i] == p[j])

	if j+1 < len(p) && p[j+1] == '*' {
		return (match && dp(s, p, i+1, j)) || dp(s, p, i, j+2)
	}

	return match && dp(s, p, i+1, j+1)
}

func isMatch(s string, p string) bool {
	return dp(s, p, 0, 0)
}

func main() {
	str := "aa"
	pattern := "a"
	ok := isMatch(str, pattern)
	fmt.Printf("is match: %v", ok)
}
