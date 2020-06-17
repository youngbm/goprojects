package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) myAbs() float64 {
	return math.Abs(v.X + v.Y)
}

func (v *Vertex) scale() {
	v.X = math.Abs(v.X * 10)
	v.Y = math.Abs(v.Y * 10)
}

func main() {
	v := Vertex{10, 10}
	fmt.Println(v.myAbs())
	v.scale()
	fmt.Println(v)
}
