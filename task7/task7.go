package main

import (
	"log"
	"math/rand"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	data := make(map[int]int)
	for i := 0; i < 7; i++ {
		wg.Add(1)
		go WriteToMap(wg, mu, data, i)
	}
	wg.Wait()
	log.Println(data)
}

// WriteToMap конкурентная запись в мапу
func WriteToMap(wg *sync.WaitGroup, mu *sync.RWMutex, data map[int]int, index int) {
	defer wg.Done()
	mu.Lock()
	data[rand.Intn(7)] = index
	mu.Unlock()
}
