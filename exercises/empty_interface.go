package main

import (
	"fmt"
)

type Describer interface {
	Describe()
}

type Reader interface {
	/* TODO: add methods */
	PrintfYouself()
}

func (x myInt) PrintfYouself() {
	fmt.Println("You are a Int")
}

type myInt int

type Persion struct {
	name string
	age  int
}

func describe(i interface{}) {
	fmt.Printf("type: %T, Value: %v\n", i, i)
}

func assert(i interface{}) {
	s := i.(int)
	fmt.Println(s)
}

func findType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("I am type int")
	case float32:
		fmt.Println("I am type float32")
	case float64:
		fmt.Println("I am type float64")
	default:
		fmt.Println("Bad type!!")
	}
}

func (p Persion) Describe() {
	fmt.Printf("I am %s,p.name, p.age %+v years old\n", p.name, p.age)
}

func findType2(i interface{}) {
	switch t := i.(type) {
	case Describer:
		t.Describe()
	case Reader:
		t.PrintfYouself()
	default:
		fmt.Println("Bad guy!")
	}

}

func main() {
	var x int = 1
	var y float32 = 2
	var z float64 = 3
	fmt.Println("Describe")
	describe(x)
	describe(y)
	describe(z)

	fmt.Println("Assert")
	assert(x)
	//assert(y)
	//assert(z)

	fmt.Println("findType")
	findType(x)
	findType(y)
	findType(z)

	p := Persion{
		"xzy",
		31,
	}

	findType2("sdfsdf")
	var fff myInt
	findType2(fff)
	findType2(p)
}
