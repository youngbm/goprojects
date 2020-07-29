package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	x := 0
	//fmt.Println(a[x++]) // x++ 是独立的表达式, 不能这样用
	x++
	fmt.Println(a[x]) //

	fmt.Println(x)
	{ //  独立的空间 ?  关注作用域
		fmt.Println(x)
		i, x := 2, 2 //  := 短变量  局部变量
		fmt.Println(i, x)
	}
	fmt.Println(x) // 1
}
