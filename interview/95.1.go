package main

import (
	"fmt"
	"time"
)

func A() int {
	time.Sleep(100 * time.Millisecond)
	return 1
}

func B() int {
	time.Sleep(1000 * time.Millisecond)
	return 2
}

func main() {
	ch := make(chan int, 1)
	go func() {
		select {
		case ch <- A():
			fmt.Println("a")
		case ch <- B():
			fmt.Println("b")
		default:
			ch <- 3
		}
	}()
	fmt.Println(<-ch)
}
