package main

import (
	"fmt"
	"time"
)

func input(c chan string) {
	for {
		c <- "from Service1111111111111111111111111111111111111"
		time.Sleep(100 * time.Millisecond)
	}
}

func output(c chan string) {
	for {
		c <- "from Service2222222222222222222222222222222222222"
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Main")
	chan1, chan2 := make(chan string, 10), make(chan string, 10)
	go input(chan1)
	go output(chan2)
	for {
		select {
		case v1 := <-chan1:
			fmt.Println(v1)
		case v2 := <-chan2:
			fmt.Println(v2)
		}
	}
}
