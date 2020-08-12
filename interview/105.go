package main

import (
	"fmt"
	"sync"
	"time"
)

var c = make(chan int)
var a int

func f() {
	a = 1
	time.Sleep(1 * time.Second)
	<-c
}

type myMutex struct {
	count int
	sync.Mutex
}

func main() {
	fmt.Println("--main-start--")
	go f()
	c <- 0 //  这里阻塞,等协程里的<-c
	fmt.Println(a)

	// Next
	var mu myMutex
	mu.Lock()
	mu.count++
	//var mu1 = mu  // 这里还在上锁，会把锁的状态也复制
	mu.Unlock()

	var mu1 = mu
	mu1.Lock()
	mu1.count++
	mu1.Unlock()
	fmt.Println(mu.count, mu1.count)
}
