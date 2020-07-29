package main

import "fmt"

func test(x int) (func(), func()) {
	return func() {
			fmt.Println(x)
			x++
		}, func() {
			fmt.Println(x) // 闭包
		}
}

func main() {
	fmt.Println("--main-start--")
	i := make([]int, 0)
	i = append(i, 1) // 扩容
	// Next

	a, b := test(100)
	a() // 100
	b() // 101
}
