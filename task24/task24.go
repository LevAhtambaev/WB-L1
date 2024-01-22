package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(10, 10)

	distance := p1.DistanceBetween(p2)
	fmt.Printf("Расстояние между точками: %f", distance)
}

// Point - структура точки
type Point struct {
	X float64
	Y float64
}

// NewPoint конструктор структуры Point
func NewPoint(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

// DistanceBetween вычисляет расстояние между двумя точками
func (p *Point) DistanceBetween(q *Point) float64 {
	dx := p.X - q.X
	dy := p.Y - q.Y
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}
