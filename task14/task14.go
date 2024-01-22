package main

import (
	"fmt"
	"reflect"
)

func main() {
	var variable interface{}

	variable = 5
	GetType(variable)

	variable = "string"
	GetType(variable)

	variable = false
	GetType(variable)

	variable = make(chan int)
	GetType(variable)
}

func GetType(variable interface{}) {
	fmt.Printf("Значение переменной: %v\nТип переменной: %s\n\n", variable, reflect.TypeOf(variable))
}
