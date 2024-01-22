package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abcd"
	fmt.Println(s1, ":", UniqueSymbols(s1)) // true

	s2 := "abCdefAaf"
	fmt.Println(s2, ":", UniqueSymbols(s2)) // false

	s3 := "aabcd"
	fmt.Println(s3, ":", UniqueSymbols(s3)) // false
}

// UniqueSymbols проверяет, что все символы в строке уникальные
func UniqueSymbols(str string) bool {
	seen := make(map[rune]bool)

	str = strings.ToLower(str)

	for _, char := range str {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}
