package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("--main-start--")
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		//wg.Add(1) // 这里直接panic
	}()

	wg.Wait()
}
