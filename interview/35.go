package main

import "fmt"

func main() {
	var b bool
	b = true
	//b = 1
	//bool(0) //不能转换
	fmt.Println(b, (1 == 2))

	var i = 10
	i++
	i--
	//var j =  i++  //  i++ 是表达式 不能赋值
	fmt.Println(i)
}
