package main

import (
	"fmt"
	"runtime/debug"
)

func myrecover() {
	if r := recover(); r != nil {
		fmt.Println("Recover from", r)
		debug.PrintStack()
	}
}

func a(x []int) {
	defer myrecover()
	fmt.Println(x[1])
}

func main() {
	debug.PrintStack()
	fmt.Println("vim-go")
	var x = []int{100}
	a(x)
	fmt.Println("This is the end of main function!!!!")
}
