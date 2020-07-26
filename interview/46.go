package main

import "fmt"

// 常量未使用是能编译通过的
const (
	x uint16 = 120
	y
	s = "abc"
	z
	a = iota // 重置为索引了  a = 4
	b = 3333
	c // 值紧跟前面
	d
	e = iota // 重置为index
	f
	jj
	yy
)

func main() {
	const a int = 123          //  可以去掉int也行
	const b float64 = 123.9123 // 会覆盖全局静态变量
	fmt.Println(a, b)

	// Next
	fmt.Println(x, y, s, z)
	fmt.Println(a, b, c, d)
	fmt.Println(e, f)
}
