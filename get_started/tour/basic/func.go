package main

import (
	"fmt"
)

func test(a, b int) int {
	return a + b
}

func combine(fn func(int, int) int, x int, y int) int {
	return fn(x, y)
}

func main() {
	//aa := test(10, 10)
	bb := func(a, b int) int {
		return a * b
	}

	fmt.Println(bb(100, 100))
	fmt.Println(combine(bb, 1000, 1000))
}
