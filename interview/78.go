package main

import (
	"fmt"
	_ "math"
)

func f()  { print("a"); panic(nil) }
func ff() { print("a"); panic(nil) }

func main() {
	fmt.Println("--main-start--")
	// 单个 case 中，可以出现多个结果选项
	// 条件表达式bu必为常量或者整数

	// Next
	//在Go代码中，注释除外，如果一个代码行的最后一个语法词段（token）为下列所示之一，
	//则一个分号将自动插入在此字段后（即行尾）：
	//1. 一个标识符；
	//2. 一个整数、浮点数、虚部、码点或者字符串字面量；
	//3. 这几个跳转关键字之一：break、continue、fallthrough和return；
	//4. 自增运算符++或者自减运算符--；
	//5. 一个右括号：)、]或}。
	//6. 为了允许一条复杂语句完全显示在一个代码行中，分号可能被插入在一个右小括号)或者右大括号}之前。

	/*
	   var (a int; b string)
	   const (M = iota; N)
	   type (MyInt int; T struct{x bool; y int32})
	   type I interface{m1(int) int; m2() string}
	*/
	/*
			var (a int; b string;);
		    const (M = iota; N;);
		    type (MyInt int; T struct{x bool; y int32;};);
		    type I interface{m1(int) int; m2() string;};

	*/
	// Next

	if x := alwaysFalse(); !x {
		//// ...
	}

	//switch alwaysFalse() // 插入; 还有后面输出true 而非false
	//{
	switch alwaysFalse(); {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	// Next
	//逗号,从不会被自动插入
	var f2 func(a int, b string) (x bool, y int)
	var f3 func(a int, b string, //// 最后一个逗号是必需的
	) (x bool, y int, //// 最后一个逗号是必需的
	)
	var _ = []int{2, 3, 5, 7, 9} //// 最后一个逗号是可选的
	//var _ = []int{2, 3, 5, 7, 9,  //// 最后一个逗号是必需的
	//}

}

func alwaysFalse() bool {
	return false
}
