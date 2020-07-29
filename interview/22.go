package main

import "fmt"

var a bool = true

func main() {
	//var x string = nil //nil 只能赋给指针类型的  arr  slice  map  chan interface
	var x string = "abc"
	if x == "abc" {
		x = "default"
	}
	fmt.Println(x)

	// Next
	defer func() {
		fmt.Println("1")
	}()

	if a == true {
		fmt.Println("2")
		return
	}

	defer func() { //没有压栈, 不会输出
		fmt.Println("3")
	}()
}
