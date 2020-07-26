package main

import "fmt"

type foo struct {
	bar int
}

func main() {
	var f foo
	var tmp int
	//f.bar, abc  := 1, 2 //:= 操作符不能用于结构体字段赋值。
	f.bar, tmp = 1, 2
	_ = tmp

	// Next
	//fmt.Println(~2)  //  这个是乜hi耶  //~ 作为按位取反运算符，Go 里面采用的是^
	z := 3                                  // 10
	fmt.Printf("%b %d %b %d", z, z, ^z, ^z) // 01 也会操作符号
	//按位取反之后返回一个每个 bit 位都取反的数，
	//对于有符号的整数来说，是按照补码进行取反操作的（快速计算方法：对数 a 取反，结果为 -(a+1) ），
	//对于无符号整数来说就是按位取反)

	var a int8 = 3
	var b uint8 = 3
	var c int8 = -3

	fmt.Printf("^%b=%b %d\n", a, ^a, ^a) //// ^11=-100 -4
	fmt.Printf("^%b=%b %d\n", b, ^b, ^b) //// ^11=11111100 252
	fmt.Printf("^%b=%b %d\n", c, ^c, ^c) //// ^-11=10 2

	// |  其中一个为1则1 与 &^ 一个0 一个1 为1
}
