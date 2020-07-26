package main

import "fmt"

type T struct {
	n int
}

func main() {
	var fn1 = func() {}
	var fn2 = func() {}

	//if fn1 == fn2 { // 函数只能和nil比较
	if fn1 != nil && fn2 != nil { // 函数只能和nil比较
		fmt.Println("equal")
	} else {
		fmt.Println("No equal")
	}

	// Next
	m := make(map[int]T)
	// m[0].n = 123 // struct 不能直接寻址
	m[0] = T{123}
	fmt.Println(m[0].n)

	x := m[0]
	fmt.Println(x.n)
}
