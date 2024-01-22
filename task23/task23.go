package main

import "log"

func main() {
	data := []int{10, 4, 8, 11, 9, 45, 24}
	index := 3

	data = Remove(data, index)
	log.Println(data)
}

// Remove удаляет i-ый элемент слайса и сохраняет порядок
func Remove(data []int, index int) []int {
	return append(data[:index], data[index+1:]...)
}
