package main

import "fmt"

func main() {
	fmt.Println("--main-start--")
	//
	var k = 9
	for k = range *[]int{} {
	}
	fmt.Println(k) // k会便覆盖,但是循环0次 k还是9
}
