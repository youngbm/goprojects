package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {
	a, b := 2, 4
	fmt.Println(add(a, b))
}
