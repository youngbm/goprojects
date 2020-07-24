package main

import "fmt"

type Integer int

func (a Integer) Add(b Integer) Integer {
	return a + b
}

//func (a *Integer) Add(b Integer) Integer { // 能改值
//	return *a + b
//}

func main() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b) //  *Integer ->  Integer 没问题
	fmt.Println(sum)
}
