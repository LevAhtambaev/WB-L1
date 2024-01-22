package main

import "fmt"

func main() {
	a := 5
	b := 10

	fmt.Printf("Изначально: a = %v, b = %v", a, b)

	a, b = b, a

	fmt.Printf("\nПоменяли местами: a = %v, b = %v", a, b)
}
