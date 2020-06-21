package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	var n = 1
	f := getfblN()
	for i := 0; i < n; i++ {
		fmt.Println(f(i))
	}

	c1 := []student{
		{
			"xzy",
			true,
		},
		{
			"fjh",
			false,
		},
		{
			"xzy",
			true,
		},
	}

	myf := func(c []student, f func(student) bool) []student {
		var rc []student
		for _, v := range c {
			if f(v) {
				rc = append(rc, v)
			}
		}
		return rc
	}
	fmt.Println(myf(c1, filter))
}

func getfblN() func(int) int {
	j, sum := 1, 1
	return func(i int) int {
		if i <= 1 {
			return 0
		}
		j, sum = sum, j+sum
		return sum
	}
}

type student struct {
	name string
	sex  bool
}

func filter(s student) bool {
	if s.sex {
		return false
	}
	return true
}
