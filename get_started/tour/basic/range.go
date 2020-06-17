package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range a {
		fmt.Println(i, v)
	}

	for _, v := range a {
		fmt.Println(v * v)
	}
}
