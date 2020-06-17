/*
*  empty interface
*
 */

package main

import "fmt"

type I interface{}

func main() {
	fmt.Println("vim-go")

	var i I = 234.22
	describe(i)

	i = 1
	describe(i)

	i = "123123"
	describe(i)
}

func describe(i I) {
	fmt.Printf("%v : %T\n", i, i)
}
