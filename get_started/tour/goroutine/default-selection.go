package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("---------------tick-------------------")
		case <-boom:
			fmt.Println("-----------------------------------Boom")
			return
		default:
			time.Sleep(50 * time.Millisecond)
			fmt.Println("                .")
		}
	}
}
