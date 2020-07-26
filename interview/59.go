package main

import (
	"fmt"
)

type N int

//func (n N) test() {  这样也可以, 但是不能改值
func (n *N) test() {
	fmt.Println(*n)
}

func main() {
	var n N = 10
	p := &n

	n++
	f1 := n.test // 这里也行 ?

	n++
	f2 := p.test

	n++
	fmt.Println(n)

	f1()
	f2()

	// Next
	var m map[int]bool
	_ = m[123]

	var pp *[5]int
	for range pp {
		_ = len(pp)
	}

	var s []int
	_ = s[:]
	//s, s[0] = []int{1, 2}, 9 // s为nil, 找不到s[0] panic
}
