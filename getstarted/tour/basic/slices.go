package main

import (
	"fmt"
)

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a[1:4])
	b := a[0:7]
	b[0] = 100
	fmt.Println(b)
	fmt.Println(a[0:2])
}
