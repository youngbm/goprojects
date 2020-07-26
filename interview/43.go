package main

import "fmt"

func main() {

	x := interface{}(nil)
	y := (*int)(nil)
	a := x == y
	b := y == nil

	//i.(Type)，其中 i 是接口，Type 是类型或接口。
	//编译时会自动检测 i 的动态类型与 Type 是否一致。但是，如果动态类型不存在，则断言总是失败。
	_, c := x.(interface{}) //  false
	fmt.Println(a, b, c)    // f t f

	x = interface{}("abc123") // interface 为空指针,可以任意赋值
	_, c = x.(interface{})
	fmt.Println(c) //  true

	// Next
	var s []int //  可以使用, 会动态扩展
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)

	//var m map[string]int  // map 不会自动扩容
	m := make(map[string]int)
	m["abc"] = 123
	fmt.Println(m)

}
