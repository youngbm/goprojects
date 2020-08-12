package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var count int

func main() {
	fmt.Println("--main-start--")
	go A()
	time.Sleep(2 * time.Second)
	mu.RLock()
	//mu.Lock()
	defer mu.RUnlock() // 读锁可以共享读，但是不能共享写
	//defer mu.Unlock()
	count++
	fmt.Println(count)
}

func A() {
	mu.RLock()
	defer mu.RUnlock()
	B()
}

func B() {
	time.Sleep(5 * time.Second)
	C()
}

func C() {
	mu.RLock()
	defer mu.RUnlock()
}
