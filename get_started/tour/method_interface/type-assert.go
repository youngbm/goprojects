package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	var i interface{} = 1234567890
	a, ok := i.(float64)
	fmt.Println(a, ok)

	b, ok := i.(int)
	fmt.Println(b, ok)
}
