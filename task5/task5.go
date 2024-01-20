package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	ch := make(chan string) // Канал для записи
	stop := make(chan bool) // Канал для послания сигнала об закрытии
	N := 5                  // Время работы программы
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go Writer(wg, ch, stop) // Запускаем отправку значений в канал в отдельной горутине
	wg.Add(1)
	go Reader(wg, ch) // Запускаем чтение значений из канала в отдельной горутине

	time.Sleep(time.Second * time.Duration(N)) // Ждем N секунд
	stop <- true
	log.Println("Channel closec")
	wg.Wait()
}

// Writer записывает в канал
func Writer(wg *sync.WaitGroup, ch chan<- string, stop <-chan bool) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			close(ch)
			return
		default:
			ch <- "message"
			time.Sleep(1 * time.Second)
		}
	}
}

// Reader читает из канала
func Reader(wg *sync.WaitGroup, ch <-chan string) {
	for val := range ch {
		log.Printf("Got %v from channel", val)
	}
	wg.Done()
}
