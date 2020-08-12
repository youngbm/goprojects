package main

import "fmt"

func main() {
	fmt.Println("--main-start--")

	a := 1
	for i := 0; i < 5; i++ {
		a = a + 1 // 用的是循环外的变量
		//a := a + 1 // 短变量, 循环类生效
		a = a * 2
	}
	fmt.Println(a) // 循环前的 a

	// Next
	result := test(10)
	fmt.Println(result)
}

func test(i int) (r int) {
	r = i * 2
	if r > 10 {
		r = 10
		//r := 10
		//短变量导致下一行return获取到这个r，而不是声明行的r
		return
	}
	return
}
