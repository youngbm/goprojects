package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(Vertex{1, 3})

	v := Vertex{123, 234}
	fmt.Println(v.x)
	fmt.Println(v.y)

	p := &v
	p.x = 100
	fmt.Println(v.x)
	fmt.Println(v.y)
}
