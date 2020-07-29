package main

import "fmt"

func main() {
	m := make(map[string]int, 2) // 这里 map2没啥意思
	m["a"] = 0
	// cap(m)  // 不能cap
	fmt.Println(len(m))

	// Next
	var x *int = nil
	//var x = nil //  x 没类型不能被nil赋值
	_ = x

	var y int
	_ = y
}
