package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v1 := Vertex{}
	v2 := Vertex{X: 1}
	v3 := Vertex{11, 22}
	p := &Vertex{111, 222}

	fmt.Println(v1, v2, v3, p.X, p.Y)
}
