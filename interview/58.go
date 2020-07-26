package main

import "fmt"

type T struct {
	x int
	y *int
}

func main() {
	i := 20
	t := T{20, &i}

	p := &t.x

	*p++
	*p--
	*p--

	t.y = p
	fmt.Println(*t.y)

	// Next
	x := make([]int, 2, 10)
	a := x[6:10] // 6,7,8,9  不包含 index 为10的
	fmt.Println(a)

	//b := x[6:] // [i:j:k] [起始:结尾:容量] 这里省略了.j默认值为长度 k默认值为容量, 转化为x[6:2:10], 则报错
	fmt.Println(b)
}
