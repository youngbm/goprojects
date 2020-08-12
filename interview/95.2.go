package main

import (
	"fmt"
	"time"
)

func service1(ch chan int) {
	ch <- 1
}
func service2(ch chan int) {
	ch <- 2
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go service1(ch1)
	go service2(ch2)
	time.Sleep(1 * time.Second)
	var n int
	select {
	case n = <-ch1:
	case n = <-ch2:
	default:
		n = 3
	}
	fmt.Println(n)
}
