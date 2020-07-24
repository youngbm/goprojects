package main

import (
	"fmt"
	"time"
)

func Foo(x interface{}) { //  x interface{}  空指针接受任何类型变量
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("Non-empty  interface")

}

func main() {
	ch := make(chan int)
	fmt.Println(len(ch), cap(ch)) // 无缓冲区的chan, 空间默认是0

	//Next
	var i *int = nil
	Foo(i)

	var j interface{} // 空指针
	Foo(j)
	j = 123
	Foo(j)
	fmt.Println(j)

	// Next
	fmt.Println("Q3")
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a", a)
		}
	}()
	fmt.Println("ok")
	time.Sleep(2 * time.Second)
	close(ch)
}
