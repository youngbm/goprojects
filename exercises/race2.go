package main

import (
	"fmt"
	"sync"
)

var count = 0

func race(m *sync.Mutex, w *sync.WaitGroup) {
	//func race(m *sync.Mutex) {
	m.Lock()
	count++
	m.Unlock()
	w.Done()
}

func main() {
	fmt.Println("main")
	var m sync.Mutex
	var w sync.WaitGroup
	for i := 0; i < 10; i++ {
		w.Add(1)
		go race(&m, &w)
		//go race(&m)
	}
	w.Wait()
	//time.Sleep(1 * time.Millisecond)
	fmt.Println(count)
}
