package main

import (
	"fmt"
)

type ct struct {
	X int
	Y int
}

func main() {
	var a [10]int
	fmt.Println(a)

	var b [10]string
	fmt.Println(b)

	b[0] = "I"
	b[1] = "am"
	b[2] = "so"
	b[3] = "handsome"
	b[4] = "!"
	fmt.Println(b)

	var c [10]ct
	fmt.Println(c)

	d := [4]int{0, 1, 2, 3}
	fmt.Println(d)
}
