package main

import (
	"fmt"
	"time"
)

func hello(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("vim-go")
	go hello("World")
	hello("Hello")
}
