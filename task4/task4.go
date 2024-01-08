package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	// Канал, в куда мы будем записывать данные
	channel := make(chan string)
	// Запуск горутины для записи в канал
	go WriteToChannel(channel)

	// Ввод числа воркеров
	var numWorkers int
	fmt.Printf("Кол-во воркеров: %v", numWorkers)
	_, err := fmt.Scan(&numWorkers)
	if err != nil {
		log.Printf("error at fmt.Scan: %v", err)
	}

	wg := &sync.WaitGroup{}

	// Запуск горутин
	for id := 0; id < numWorkers; id++ {
		wg.Add(1)
		go Worker(wg, channel, id)
	}
	// По нажатию ctrl c программа завершится
	s := <-c
	fmt.Println("Got signal:", s)
	// Закрываем канал, ждем пока все горутины завершат работу
	close(channel)
	wg.Wait()
}

// WriteToChannel Функция для записи данных в канал
func WriteToChannel(ch chan<- string) {
	for {
		// Записали данные в канал
		ch <- "сообщение"
		time.Sleep(1 * time.Second)
	}
}

// Worker Воркер будет читать с канала и выводить данные
func Worker(wg *sync.WaitGroup, ch <-chan string, id int) {
	defer wg.Done()
	for info := range ch {
		fmt.Printf("worker №%d: %s\n", id, info)
	}
}
