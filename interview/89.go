package main

import "fmt"

type P *int
type Q *int

func main() {
	fmt.Println("--main-start--")

	v := []int{1, 2, 3}
	for i, n := 0, len(v); i < n; i++ {
		v = append(v, i)
	}
	fmt.Println(v)

	// Next
	var p P = new(int)
	*p += 8
	var x *int = p
	var q Q = x
	*q++

	fmt.Println(*p, *q)
}
