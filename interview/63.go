package main

import "fmt"

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func main() {
	fmt.Println("--main-start--")
	f := F(5)

	defer func() {
		fmt.Println(f()) // 闭包N.O 4 8-->9
	}()

	defer fmt.Println(f()) //  N.O 1  5 --> 6  压栈,先处理括号里面触发执行,执行第一次 6

	defer f() //N.O 3 7-->8 压栈，但是不触发执行,返回的时候才执行

	i := f()       // N.0 2 6 --> 7  执行第二次
	fmt.Println(i) // 7
}
