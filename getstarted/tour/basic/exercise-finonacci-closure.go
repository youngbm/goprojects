package main

import "fmt"

func finanocci() func(int) int {
	ib := 0
	ret := 0
	return func(i int) int {
		if i <= 1 {
			ret += i
			return ret
		} else {
			ret, ib = ib+ret, ret
			return ret
		}
	}
}

func main() {
	fmt.Println()
	a := finanocci()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
}
