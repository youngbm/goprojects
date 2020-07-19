package main

import "fmt"

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) // 10, 1, 2  3\n1, 1, 3, 4  defer 里面调用其他函数,会立即生效

	a = 0
	defer calc("2", a, calc("20", a, b)) // 20, 0, 2  2\n2, 0, 2, 2  defer 里面调用其他函数,会立即生效
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// 10, 1, 2  3
// 20, 0, 2  2
// 2, 0, 2, 2
// 1, 1, 3, 4
//range  map  出来的是没有顺序的
