package main

import "fmt"

type I interface {
	printAll()
}

type myFloat float64

func (f myFloat) printAll() {
	fmt.Println(f)
}

type Vertex struct {
	s string
}

func (s Vertex) printAll() {
	fmt.Println(s)
}

func main() {
	f := myFloat(123123)
	v := Vertex{"mystring"}
	var i I = f
	i.printAll()
	describe(i)
	fmt.Printf("\n\n")
	i = v
	i.printAll()
	describe(i)
}

func describe(i I) {
	fmt.Printf("print something: %v: %T \n", i, i)
}
