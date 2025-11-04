package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool{
	s = strings.ToLower(s)

	var removed string
	for _, ch := range s{
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			removed += string(ch)
		
	}
}

n := len(removed)
	for i := 0; i < n/2; i++ {
		if removed[i] != removed[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	st := []string{
		"New Year",
		"noon",
		"12321",
	}

	for _, text := range st {
		fmt.Printf("\"%s\" -> %v\n", text, IsPalindrome(text))
	}
}
