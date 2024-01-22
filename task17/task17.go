package main

import (
	"fmt"
	"log"
)

func main() {
	items := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	elem := 12

	index, err := BinarySearch(items, elem)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Elem: %v, Position: %v", elem, index)
}

// BinarySearch выполняет бинарный поиск нужного элемента
func BinarySearch(items []int, elem int) (int, error) {
	if len(items) == 1 {
		if items[0] != elem {
			return 0, fmt.Errorf("elem %v not found", elem)
		}
		return 0, nil
	}
	if elem >= items[len(items)/2] {
		position, err := BinarySearch(items[len(items)/2:], elem)
		if err != nil {
			return 0, err
		}
		return len(items)/2 + position, nil
	} else {
		position, err := BinarySearch(items[:len(items)/2], elem)
		if err != nil {
			return 0, err
		}
		return position, nil
	}
}
