package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Before sleep")
	Sleep(2 * time.Second)
	fmt.Println("After sleep")
}

// Sleep реализован через time.Timer
func Sleep(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
}
