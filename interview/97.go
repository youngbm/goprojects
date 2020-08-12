package main

import (
	"fmt"
)

func changeMap(m map[string]int) {
	m["a"] = 100
}

func main() {
	fmt.Println("--main-start--")
	m := make(map[string]int, 100)
	m["a"] = 1
	fmt.Println(m)
	// Map是一种引用类型
	// Map是对底层数据的引用。
	// 编写代码的过程中，会涉及到Map拷贝、函数间传递Map等。
	// 跟Slice类似，Map指向的底层数据是不会发生copy的。
	changeMap(m)
	fmt.Println(m)
}
