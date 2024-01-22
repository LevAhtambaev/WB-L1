package main

import "log"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	res := MakeSet(words)
	log.Println(res)
}

// MakeSet возвращает множество из набора слов
func MakeSet(words []string) []string {
	set := make(map[string]bool)
	for _, str := range words {
		set[str] = true
	}
	var res []string
	for key := range set {
		res = append(res, key)
	}
	return res
}
