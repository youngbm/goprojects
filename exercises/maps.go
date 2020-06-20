package main

import "fmt"

type i struct {
	X int
	Y int
}

type niub struct {
	a []int
	b i
}

func main() {
	var j = i{
		1,
		2,
	}
	fmt.Println(j)
	var x = make(map[string]int)
	x["a"] = 1
	x["b"] = 2
	fmt.Println(x["a"])

	var xzy = make([]niub, 2)
	xzy = []niub{
		niub{
			[]int{1, 2, 3},
			i{
				100,
				200,
			},
		},
		niub{
			[]int{1, 2, 3},
			i{
				100,
				200,
			},
		},
	}
	fmt.Println(xzy)
}
