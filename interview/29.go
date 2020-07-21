package main

import (
	"fmt"
	"time"
)

type abc struct {
	x, y, z int
}

var a = abc{x: 1}
var b = abc{x: 1}

//var a = &abc{x: 1}
//var b = &abc{x: 1}

func main() {
	if a == b {
		fmt.Println(true)
	} else {
		fmt.Println("false")
	}

	// Next
	v := []int{1, 2, 4}
	for i := range v { // range 只有一个接受变量，则为index
		v = append(v, i)
	}

	// Next
	var m = []int{1, 2, 3, 4, 5, 6, 9}
	for k, v := range m {
		go myPrint(k, v) // 正常打屏
		// 闭包,异步调用. 启动时 k,v => 6,9
		// var k int = k //  声明新的变量, 正常打印
		// var v int = v
		k := k //  声明新的变量, 正常打印
		v := v
		go func() {
			fmt.Println(k, v)
		}()
	}
	time.Sleep(1 * time.Second)
}

func myPrint(k, v int) {
	fmt.Println(k, v)
}
