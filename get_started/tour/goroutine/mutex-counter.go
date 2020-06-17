package main

import (
	"fmt"
	"sync"
)

type saveCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *saveCounter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *saveCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	fmt.Println("vim-go")
	s := saveCounter{v: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		s.Inc("somekey")
	}
	fmt.Println(s.Value("somekey"))
}
