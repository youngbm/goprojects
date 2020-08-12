package main

import (
	"fmt"
	"time"
)

var ch chan int = make(chan int)

func generate() {
	for i := 12; i < 5000; i += 17 {
		ch <- i
		time.Sleep(1 * time.Microsecond)
	}
}

func main() {
	fmt.Println("--main-start--")
	timeout := time.After(800 * time.Millisecond)
	go generate()
MAIN_LOOP: // 这里报错会自动缩进 vim-go 太2了
	found := 0
	for {
		select {
		case v1, ok := <-ch:
			if ok {
				if v1%38 == 0 {
					fmt.Println(v1, "is a multiple of 17 and 38")
					found++
					if found == 3 {
						break MAIN_LOOP // 跳到要跳出的循环中
					}
				} else {
					break MAIN_LOOP
				}
			}
		case <-timeout:
			fmt.Println("time out")
			break MAIN_LOOP
		}

	}
	fmt.Println("The end!!!")
}
