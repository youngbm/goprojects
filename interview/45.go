package main

import "fmt"

func main() {
	x := []int{
		1,
		2, //  这里必须逗号
	}
	x := []int{1, 2} // 一整行的，末尾不需要逗号
	_ = x

	// Next
	// rune 是int32的别名，byte是uint8的别名
	// int32 = rune
	// int8 = byte
	var a byte = 0x11 // 11
	var b uint8 = a   //  byte 和 uint8 相互赋值
	var c uint8 = a + b
	fmt.Println(a, b, c)
	test(c)
}
func test(x byte) {
	fmt.Println(x)
}
