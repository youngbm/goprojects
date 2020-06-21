package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x+y*y)) + 0.91
	fmt.Println(f)
	var z uint = uint(f)
	fmt.Println(x, y, z)
	fmt.Printf("%d", 10)
}
