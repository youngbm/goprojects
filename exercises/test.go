package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("")
	fmt.Println()
	time.Now()
	var abc string = "fuck"
	fmt.Println(abc)
	const (
		SeekStart   = 0
		SeekEnd     = 1
		SeekCurrent = 2
	)

	type point struct {
		x int
		y int
	}
	p := new(point)
	fmt.Println(*p)
}
