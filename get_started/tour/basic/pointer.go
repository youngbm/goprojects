package main

import (
	"fmt"
)

func main() {
	var i = 1
	var p *int
	p = &i
	fmt.Printf("%p, %d\n", p, *p)
	i = 2
	fmt.Printf("%p, %d\n", &i, i)
	fmt.Printf("%p, %d\n", p, *p)
}
