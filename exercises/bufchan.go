package main

import (
	"fmt"
	"time"
)

func writeChan(ch chan int) {
	for i := 0; i < 90; i++ {
		time.Sleep(100 * time.Microsecond)
		ch <- 3
	}
}

func main() {
	fmt.Println("--main-start--")
	var ch = make(chan int)
	go writeChan(ch)
	for i := 0; i < 40; i++ {
		time.Sleep(1 * time.Second)
		select {
		case ch <- 2:
			fmt.Println("write")
		case v := <-ch:
			fmt.Println("read", v)
		default:
			fmt.Println("default")
		}
	}
}
