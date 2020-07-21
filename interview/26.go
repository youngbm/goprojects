package main

import "fmt"

const (
	a = iota
	b = iota
)

const (
	c = "name" // 后面iota 则为1
	d = iota   // 前面有数 则从1开始
	e = iota
)

type People interface {
	Show()
}

type Student struct{}

func (s Student) Show() {
}

func main() {
	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c)
	fmt.Println(d) //  1
	fmt.Println(e) //  2

	// Next
	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not  nil")
	}

	// 当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil
	var p People = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}
