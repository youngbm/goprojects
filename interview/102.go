package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var chain string

func A() {
	mu.Lock
	defer mu.Unlock
	chain = chain + " --> A"
	B() // 锁还没释放就又上锁了
}

func B() {
	chain +=  " --> B"
	C()
}

func C() {
	mu.Lock
	defer mu.Unlock
	chain = chain + " --> C"
}

func main() {
	fmt.Println("--main-start--")
	chain =  "main"
	A()
	fmt.Println(chain)
}
