package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
)

func main() {
	var method int
	fmt.Print("Введите номер способа: ")
	_, err := fmt.Scan(&method)
	if err != nil {
		log.Printf("Error at scan: %v", err)
	}
	wg := &sync.WaitGroup{}
	ch := make(chan string)
	switch method {
	case 1:
		wg.Add(1)
		done := make(chan bool)
		go SignalChannel(wg, done)
		done <- true
	case 2:
		wg.Add(1)
		go ChannelClose(wg, ch)
		close(ch)
	case 3:
		ctx, cancel := context.WithCancel(context.Background())
		wg.Add(1)
		go Context(ctx, wg)
		cancel()
	case 4:
		wg.Add(1)
		go GoExit(wg)
	default:
		fmt.Println("Выберите из 4 доступных")
	}
	wg.Wait()

}

// SignalChannel используем канал в качестве сигнала для закрытия
func SignalChannel(wg *sync.WaitGroup, done chan bool) {
	for {
		select {
		case <-done:
			log.Println("goroutine is ending")
			wg.Done()
			return
		default:
			// основные действия
		}
	}
}

// ChannelClose через закрытие канала
func ChannelClose(wg *sync.WaitGroup, ch chan string) {
	for value := range ch {
		log.Println(value)
	}
	log.Println("goroutine is ending")
	wg.Done()
}

// Context через контекст
func Context(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			log.Println("goroutine is ending")
			wg.Done()
			return
		default:
			// основные действия
		}
	}
}

// GoExit через runtime.GoExit
func GoExit(wg *sync.WaitGroup) {
	for {
		// основные действия
		log.Println("goroutine is ending")
		wg.Done()
		runtime.Goexit()
	}
}
