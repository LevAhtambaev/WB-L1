package main

import "fmt"

type Human struct { //Human с произвольным набором полей
	Field1 string
	Field2 string
}

// Method1 первый метод структуры Human
func (h *Human) Method1() {
	fmt.Println("Method1")
}

// Method2 первый метод структуры Human
func (h *Human) Method2() {
	fmt.Println("Method2")
}

type Action struct {
	Human         // Human как часть структуры Action даст возможность вызывать методы Human
	Field1 string // Могут быть любые произвольные поля
	Field2 string
}

func main() {
	action := Action{}
	action.Method1() // Method1 доступен из структуры Action
	action.Method2() // Method2 доступен из структуры Action
}
