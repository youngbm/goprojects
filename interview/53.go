package main

import "fmt"

type T struct {
	n int
}

func (t *T) Set(n int) {
	t.n = n
}

func getT() T {
	return T{}
}

func main() {

	// 1. 结构体不能寻址;
	// 2. 不可寻址的结构体不能调用带结构体指针接收者的方法；
	//getT().Set()

	g := getT()
	g.Set(123)
	fmt.Println(g.n)
}
