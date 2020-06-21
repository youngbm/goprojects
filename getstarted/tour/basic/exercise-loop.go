package main

import (
	"fmt"
	"math"
)

func getSqrt(x float64, looptime int) float64 {
	var z float64
	z = 1
	for i := 0; i < looptime; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	var looptime int
	for looptime = 0; looptime < 10; looptime++ {
		fmt.Println(getSqrt(1024, looptime))
	}
	fmt.Println(math.Sqrt(1024))
}
