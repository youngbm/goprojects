package main

import (
	"fmt"
	"runtime"
)

func main() {

	s := make([]int, 1) //func(t Type, size ...IntegerType) Type
	fmt.Println(s)
	runtime.GOMAXPROCS(2)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"

	select { // 随机读取IO操作
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value) // 看先读到哪个chan, 有可能触发到panic
	}
}
