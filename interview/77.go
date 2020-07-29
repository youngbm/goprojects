package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("--main-start--")
	ch := make(chan int, 10)
	fmt.Println(cap(ch))

	// Next

	t := &testing.T{}
	Test13(t)
}

func hello(num ...int) {
	num[0] = 18
}

func Test13(t *testing.T) {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0]) // 18
}
