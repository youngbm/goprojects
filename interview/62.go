package main

import "fmt"

func main() {
	fmt.Println("--main-start--")
	nil := 123       //  nil会被覆盖
	fmt.Println(nil) // 123
	//var _ map[string]int = nil // nil 可以赋值到 map, 但是这里nil被覆盖为 int

	// Next
	var i int8 = -128
	fmt.Println(i / -1) // -128 溢出了

}
