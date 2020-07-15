package main

import "fmt"

func main() {
	fmt.Println("Main")
	//var mych chan int       // 只是声明一个chan int变量
	ch := make(chan int, 2) // 带缓冲的读写chan
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(ch) // 记得关闭 chan

	// Next
	type person struct {
		name string
	}

	var m map[person]int
	p := person{"john"}
	fmt.Println(m[p])

	// Next
	i := []int{1, 2, 3, 4, 5, 6}
	hello(i...) // i 为数组
	fmt.Println(i)
}

func hello(m ...int) {
	m[0] = 100 //  m[0] 是改地址的值不是副本，所以会修改到真实值
}
