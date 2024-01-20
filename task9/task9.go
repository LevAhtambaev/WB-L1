package main

import (
	"log"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	input := make(chan int)  // Канал для записи чисел (x)
	output := make(chan int) // Канал для записи результата операции x*2

	wg.Add(1)
	go WriterInput(wg, input) // Записываем числа (x) в канал input

	wg.Add(1)
	go WriterOutput(wg, input, output) // Записываем числа x*2 в канал output

	wg.Add(1)
	go Reader(wg, output) // Выводим данные из канала output в stdout

	wg.Wait()
}

// WriterInput записывает в канал числа из массива
func WriterInput(wg *sync.WaitGroup, input chan int) {
	defer wg.Done()
	numbers := [5]int{2, 4, 6, 8, 10} // Пример массива чисел
	for _, num := range numbers {
		input <- num
	}
	close(input)
}

// WriterOutput умножает числа из канала на 2 и записывает в другой канал
func WriterOutput(wg *sync.WaitGroup, input, output chan int) {
	defer wg.Done()
	for num := range input {
		output <- num * 2
	}
	close(output)
}

// Reader читатет из канала и выводит в stdout
func Reader(wg *sync.WaitGroup, input chan int) {
	defer wg.Done()
	for res := range input {
		log.Println(res)
	}
}
