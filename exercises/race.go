package main

import (
	"fmt"
	"sync"
)

func r(c chan int, count *int, w *sync.WaitGroup) {
	c <- 0
	*count++
	<-c
	w.Done()
}

func main() {
	fmt.Println("main")
	var c = make(chan int, 1)
	var count int = 0
	var w sync.WaitGroup
	for i := 0; i < 10000; i++ {
		w.Add(1)
		go r(c, &count, &w)
	}
	w.Wait()
	fmt.Println(count)
}
