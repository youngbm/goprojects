package main

import "fmt"

var x int

func init() {
	x++
}

func min(a int, b uint) {
	min := 0
	min = copy(make([]int, a), make([]int, b)) // 利用copy 的特性
	fmt.Printf("The min of %d and %d is %d\n", a, b, min)
}

func main() {
	fmt.Println("--main-start--")
	//init() // init 不能被任何函数调用
	fmt.Println(x)

	// Next
	min(123, 234)
}
