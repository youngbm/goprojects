package main

import (
	"fmt"
	"math"
)

func pow(a, b, c float64) float64 {
	if v := math.Pow(a, b); v < c {
		return v
	} else {
		return 1
	}
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 10),
	)
}
