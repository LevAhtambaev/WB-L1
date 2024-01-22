package main

import (
	"log"
)

func main() {
	set1 := []int{10, 11, 3, 5, 9}
	set2 := []int{11, 3, 6, 7, 8}

	result := Intersection(set1, set2)
	log.Println(result)
}

// Intersection реализация пересечения множеств
func Intersection(set1, set2 []int) []int {
	firstSet := make(map[int]bool, len(set1))
	intersection := make([]int, 0)

	for _, num := range set1 {
		firstSet[num] = true
	}

	for _, num := range set2 {
		if firstSet[num] {
			intersection = append(intersection, num)
		}
	}

	return intersection
}
