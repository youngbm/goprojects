package main

import (
	"fmt"
)

func main() {
	fmt.Println("Append.......")
	a := make([]int, 5)

	c := append(a, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(c)
	printSlices(c)

	d := []int{1, 2, 3}
	printSlices(d)

	k := append(d, 5, 6, 7, 8, 9, 0)
	printSlices(k)
}

func printSlices(l []int) {
	fmt.Printf("Len:%d; Cap: %d; V:%v\n", len(l), cap(l), l)
}
