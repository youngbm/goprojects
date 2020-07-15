package main

import (
	"fmt"

	"goprojects/interview/foo" // 包里面有2个文件,bar.go和foo.go,各个文件的init()都会被调用
	_ "goprojects/interview/foo"
)

type words struct {
	s string
}

type i interface {
	hello()
}

func (w words) hello() {
	fmt.Println("I am hello in function")
}

// 不论什么包都是先执行init
func init() {
	fmt.Println("main: I am in init")
}

func main() {
	fmt.Println("--main-start--")

	// init() // 报错 undefiend. init() 函数只能作为初始化的操作不能表其他调用
	foo.SayBar() // 如果想被调用必须要首字母大写

	// Next
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("Not nil")
		fmt.Println(h)
	}

	// Next
	//var i int = 100 //  cannot type switch on non-interface value i(type int) //var j = GetValue
	//var j []int
	var j words
	j = words{"123123123"}
	j.hello()

	var x interface{}    // 不赋值的时候, x.(type) is nil
	x = words{"6666666"} // 赋值的时候, x.(type) 是接口类型  interface{}

	//x.hello() // x是interface 不是 words 所以不能调用方法hello.
	//报错:  x.hello undefined (type interface {} is interface with no methods)

	//fmt.Println(x.(type)) // 只能在switch中
	switch x.(type) {
	//switch j.(type) {只能对接口类型操作, i.(type),并且只能在switch里面进行
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case interface{}:
		fmt.Println("interface{}")
	case func():
		fmt.Println("func(){}")
	default:
		fmt.Println("nil")
	}
}

func GetValue() int {
	return 1
}

func hello() []string {
	return nil
}
