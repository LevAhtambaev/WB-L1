package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x, y float64

	x = 13364277
	y = 21486951

	firstNum := big.NewFloat(x)
	secondNum := big.NewFloat(y)

	resultNum := new(big.Float).SetPrec(100)

	resultNum.Mul(firstNum, secondNum)
	fmt.Println("Умножение:")
	fmt.Println(resultNum)

	resultNum.Quo(firstNum, secondNum)
	fmt.Println("Деление, используя big:")
	fmt.Println(resultNum)

	resultNum.Add(firstNum, secondNum)
	fmt.Println("Сложение:")
	fmt.Println(resultNum)

	resultNum.Sub(firstNum, secondNum)
	fmt.Println("Вычитание:")
	fmt.Println(resultNum)

}
