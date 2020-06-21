package main

import (
	"fmt"
)

func sample() func(int) int {
	sum := 0
	return func(a int) int {
		sum += a
		return sum
	}
}

func main() {
	pos, neg := sample(), sample()
	for i := 0; i < 3; i++ {
		fmt.Println(pos(i))
		fmt.Println(neg(2*i + 2))
		fmt.Println("\n")
	}
}
