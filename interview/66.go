package main

import "fmt"

type T struct {
	n int
}

func main() {
	fmt.Println("--main-start--")
	ts := [2]T{}
	for i, t := range ts { // ts 为副本
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9 // 结构体不可以寻址
		case 1:
			fmt.Println(t.n, " ") //  t.n 是 0
		}
	}
	fmt.Println(ts, cap(ts))
}
