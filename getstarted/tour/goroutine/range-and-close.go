package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		time.Sleep(1000 * time.Millisecond)
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	fmt.Println("vim-go")
	c := make(chan int, 20)
	go fibonacci(cap(c), c)
	// 用range 只返回一个值
	//for i := range c {
	//	fmt.Println(i)
	//}
	for {
		i, ok := <-c
		if !ok {
			fmt.Println(ok)
			break
		}
		fmt.Println(i)
	}
}
