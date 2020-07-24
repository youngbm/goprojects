package main

import "fmt"

type T1 struct {
	s string
}

func main() {

	//var t *T1  //  只是指针变量, 引发异常
	var t T1 //  有相关的地址空间
	t.s = "abc"

	//fmt.Println(10 / 0) //division by zero panic
	s := [3]int{1, 2, 3}
	//fmt.Println(s[10]) // 下标越界 panic

	for v := range s {
		fmt.Print(v)
	}
}
