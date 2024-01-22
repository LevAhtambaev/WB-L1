package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int32
	wg    *sync.WaitGroup
}

func main() {
	counter := Counter{
		wg: &sync.WaitGroup{},
	}

	for i := 0; i < 7; i++ {
		counter.wg.Add(1)
		go counter.Increment()
	}

	counter.wg.Wait()
	fmt.Printf("Counter: %v", counter.value)
}

// Increment инкрементирование счетчика в конкурентной среде, с использованием атомарной операции
func (c *Counter) Increment() {
	atomic.AddInt32(&c.value, 1)
	c.wg.Done()
}
