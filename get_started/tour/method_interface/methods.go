package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) abs() float64 {
	return math.Abs(v.X + v.Y)
}

func main() {
	a := Vertex{100, 100}
	fmt.Println(a.abs())
}
