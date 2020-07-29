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

	var j interface{} = &a
	sum0 := j.(*Integer).Add(b) //j.(*Integer)变为和方法一样的
	fmt.Println(sum0)

	var i interface{} = a
	sum1 := i.(Integer).Add(b) //  Integer ->  Integer 没问题
	// i.Add()  i为接口,报错。i.(Interger)才是接收者
	fmt.Println(sum1)
}
