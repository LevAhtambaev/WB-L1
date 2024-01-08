package main

import (
	"fmt"
	"sync"
)

func main() {
	input := [5]int{2, 4, 6, 8, 10} // Входные данные
	wg := &sync.WaitGroup{}

	for _, number := range input { // Запуск горутин
		wg.Add(1)
		go Square(wg, number)
	}
	wg.Wait()
}

// Square подсчет квадрата числа
func Square(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	square := num * num
	fmt.Println(square)
}
