package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "snow dog sun"
	reversedSentence := GetReverseWords(sentence)
	fmt.Printf("Слова в другом порядке: %v", reversedSentence)
}

// GetReverseWords переворачивает слова в строке
func GetReverseWords(sentence string) string {
	words := strings.Split(sentence, " ")

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
