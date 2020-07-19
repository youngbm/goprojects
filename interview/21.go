package main

import "fmt"

type S struct {
}

func f(x interface{}) {
}

//[Important] interface{} 时可以接收任何类型的参数，包括用户自定义类型等
func g(x *interface{}) {
}

type SS struct {
	m string
}

func ff() *SS {
	return &SS{"Created by func SS"}
}

func main() {
	var a []int     // nil  不会分配内存，优先选择
	var b = []int{} // []
	fmt.Println(a, b)

	// Next
	s := S{}
	p := &s

	f(s)
	//g(p)
	fmt.Sprintln(p)
	f(s)
	//g(p)

	// Next
	pp := ff()
	fmt.Println(pp.m)
}
