package main

import (
	"fmt"
)

func main() {
	type Vertex struct {
		Lat, Long float64
	}

	a := make(map[string]Vertex)
	a["fuck"] = Vertex{123123.123, 12312.123}
	fmt.Println(a["fuck"].Long)
	fmt.Println(a["fuck"].Lat)

	b := map[string]Vertex{
		"fuck0": Vertex{
			123123, 122,
		},
		"fuck1": Vertex{
			1, 2,
		},
	}
	fmt.Println(b)

	c := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(c)
}
