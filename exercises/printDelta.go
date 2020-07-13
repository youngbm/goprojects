package main

import (
	"fmt"
	"time"
)

func printDelta(N int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N*2; j++ {
			if i == N-1 {
				if j == 0 || j > N*2 {
					fmt.Printf(" ")
					continue
				}
				fmt.Printf("*")
			} else if j == N-i || j == N+i {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	for i := 0; i < 10; i++ {
		printDelta(i)
		time.Sleep(300 * time.Millisecond)
	}
}
