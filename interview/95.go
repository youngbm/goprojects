package main

import "fmt"

type foo struct {
	Val int
}

type bar struct {
	Val int
}

func main() {
	fmt.Println("--main-start--")
	c := foo{Val: 5}
	d := bar{Val: 5}
	e := bar{Val: 5}         // 6 则false
	fmt.Println(c == foo(d)) // true
	fmt.Println(e == d)      // true
}
