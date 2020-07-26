package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	fmt.Println("hello world")

	var x = []int{8: 2, 100, 4: 3} //  {i:e}  i为index, e位值其余补零
	fmt.Println(x)

	// Next
	v := 1
	incr(&v)
	fmt.Println(v)

}
