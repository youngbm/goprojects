package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}
	t := a[3:4:4] // index 3-4 不包含4  这里的4毫无意义
	fmt.Println(t)
	fmt.Printf("%d-%d\n", len(t), cap(t)) //截取原来的数据, 1-1
	fmt.Println(t[0])

	b := [10]int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}
	if a == b { // [5]int和[6]int直接不能比较了, [5]int和[5]int才能比较
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	x := 5
	y := 5.1
	//fmt.Println(x + y) //  进行了类型定义，不能直接相加
	fmt.Println(x + int(y))
	fmt.Println(5 + 5.11111) // 可以直接相加, 会自己转换数据类型
}
