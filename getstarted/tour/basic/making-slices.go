package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	printSlices("a", a)

	b := make([]int, 3, 5)
	printSlices("b", b)

	c := b[:2]
	printSlices("c", c)

	d := a[2:4]
	printSlices("d", d)
}

func printSlices(s string, l []int) {
	fmt.Printf("name: %s; len: %d; cap: %d; v: %v \n", s, len(l), cap(l), l)
}
