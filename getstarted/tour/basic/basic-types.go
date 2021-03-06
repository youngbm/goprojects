package main

import (
	"fmt"
	"math/cmplx"
)

func test123() {
	fmt.Println(123)
}

func test123123() {
	fmt.Println(123)
}

func main() {

	test123()

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T, Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T, Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T, Value: %v\n", z, z)
}
