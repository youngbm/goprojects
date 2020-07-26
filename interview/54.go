package main

import "fmt"

type N int

func (n N) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {
	*n++
	fmt.Printf("v: %p, %v\n", n, *n)
}

func main() {
	var a N = 25

	p := &a
	p1 := p
	//p1 := &p  // 取值后就不能用p的方法了
	p1.value()
	p1.pointer()

	// Next
	// [Important] 方法也能象函数一样直接赋值
	a++
	f1 := N.value
	f1(a) // 调用方式牛逼, 制定方法的接收者

	f2 := (*N).value
	f2(&a) // 调用方式牛逼, 制定方法的接收者
}
