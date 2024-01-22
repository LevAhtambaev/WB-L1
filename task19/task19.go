package main

import (
	"fmt"
)

func main() {
	str := "главрыба"
	reversedStr := GetReverseString(str)
	fmt.Println(reversedStr)
}

// GetReverseString переворачивает строку
func GetReverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
