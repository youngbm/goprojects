package main

import "fmt"

var p = fmt.Print

func main() {
	ch := make(chan int, 1)

	// _,_  不是新变量 不需要 短等号:=
	for _, _ = range [9]int{1, 2} { // 循环9次啊 range 跟进容量来循环
		select {
		case <-ch:
			p(2)
			ch = nil //  读nil频道读出0
		case ch <- 0:
			p(1)
		default:
			p(0)
		}
	}

}
