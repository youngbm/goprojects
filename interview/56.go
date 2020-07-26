package main

import "fmt"

type N int

func (n N) test() {
	fmt.Println(n)
}

func (n *N) testp() {
	fmt.Println(*n)
}

func main() {
	var n N = 10
	p := &n

	// 2种调用方式都是可以的
	f1 := n.test
	f11 := N.test

	f2 := p.test
	f22 := (*N).testp // 括号括起来才行 666

	f1()
	f2()
	f11(n)
	f22(p)
}
