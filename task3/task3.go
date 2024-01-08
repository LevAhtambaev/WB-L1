package main

import (
	"fmt"
	"sync"
)

func main() {
	input := [5]int{2, 4, 6, 8, 10} // Входные данные
	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	output := new(int)
	for _, num := range input { // Запуск горутин
		wg.Add(1)
		go Square(wg, mu, num, output)
	}
	wg.Wait()
	fmt.Println(*output)
}

// Square подсчет квадрата числа и добавление его к выходному числу
func Square(wg *sync.WaitGroup, mu *sync.RWMutex, num int, output *int) {
	defer wg.Done()
	square := num * num
	mu.Lock()
	*output = *output + square
	mu.Unlock()
}
