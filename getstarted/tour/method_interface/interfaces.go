package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X float64
	Y float64
}

type myFloat float64

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.Y*v.Y + v.X*v.X)
}

func (f myFloat) Abs() float64 {
	return math.Sqrt(float64(f))
}

func main() {
	fmt.Println("vim-go")
	var i Abser

	v := Vertex{400, 300}
	f := myFloat(1.723242)

	i = &v
	fmt.Println(i.Abs())
	fmt.Println(v.Abs())

	i = f
	fmt.Println(i.Abs())
	fmt.Println(f.Abs())

	// 隐shi调用
	var j Abser = &Vertex{3, 4}
	fmt.Println(j.Abs())
	var m Abser = myFloat(4.00)
	fmt.Println(m.Abs())
}
