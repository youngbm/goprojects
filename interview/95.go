package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 1)

type foo struct {
	Val int
}

type bar struct {
	Val int
}

func A() int {
	time.Sleep(1000 * time.Millisecond)
	return 1
}

func B() int {
	time.Sleep(1000 * time.Millisecond)
	return 2
}

func main() {
	fmt.Println("--main-start--")
	c := foo{Val: 5}
	d := bar{Val: 5}
	e := bar{Val: 5}         // 6 则false
	fmt.Println(c == foo(d)) // true
	fmt.Println(e == d)      // true

	// Next
	for i := 0; i < 20; i++ {
		go func() {
			select {
			case ch <- A():
			case ch <- B():
			default:
				ch <- 3
			}
		}()
		fmt.Println(<-ch)
	}
}
