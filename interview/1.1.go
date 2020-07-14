package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start....")
	for w := 0; w < 10; w++ {
		go f1()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("end....")
}

func f1() {
	defer recover()
	panic("f1")
	go f2()
}

func f2() {
	defer recover()
	panic("f2")
}
