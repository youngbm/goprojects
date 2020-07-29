package main

import "fmt"

func change(s ...int) {
	s = append(s, 3)
}

func main() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...) //  slice 传入可变函数，不会创建新切片
	fmt.Println(slice)

	change(slice[0:3]...) // 操纵同一个底数组,创建新切片len2, cap5.会修改到底数组
	fmt.Println(slice)

	// Next
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for k, v := range a { //
		if k == 0 {
			//if k == 3 {   // 过了之后才改值
			a[1] = 12
			a[2] = 13
		}
		r[k] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}
