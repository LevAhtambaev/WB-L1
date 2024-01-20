package main

import (
	"fmt"
)

func main() {
	var num int64 = 7 // 7 = 111 в двоичной
	index := 1        // будем менять бит с индексом 1, число поменяется на 101 = 5

	fmt.Printf("Начальное значение: %d\n", num)

	res := SetBit(num, index, 0) // Установка i-го бита в 0 (101 = 5)
	fmt.Printf("Установили %d-ый бит в 0, получили: %d\n", index, res)

	res = SetBit(num, index, 1) // Установка i-го бита в 1 (111 = 7)
	fmt.Printf("Установили %d-ый бит в 1, получили: %d\n", index, res)
}

// SetBit устанавливает i-й бит в числе в 1 или 0
func SetBit(n int64, index int, bit int) int64 {
	var result int64
	mask := int64(1) << index
	if bit == 1 {
		result = mask | n
	} else {
		result = n & ^mask
	}
	return result
}
