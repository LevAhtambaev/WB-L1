package main

import (
	"log"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, -21.0, 13.0, -9.1, 5.2, 19.0, 15.5, 24.5, 32.5}

	groups := NumbersPerGroup(temps)

	log.Println(groups)
}

// NumbersPerGroup разграничивает числа по группам с шагом 10
func NumbersPerGroup(temps []float64) map[int][]float64 {
	var step = 10.0
	groups := make(map[int][]float64)
	for _, temp := range temps {
		switch {
		case temp >= 0 && temp < step:
			groups[1] = append(groups[1], temp)
		case temp < 0 && math.Abs(temp) < step:
			groups[-1] = append(groups[-1], temp)
		default:
			decimals := int(temp/step) * int(step)
			groups[decimals] = append(groups[decimals], temp)
		}
	}
	return groups
}
