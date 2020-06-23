package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	n := new(int)
	*n = 100
	fmt.Println(*n)

	type person struct {
		name string
		age  int
	}
	fmt.Println(make([]person, 10))

	goto flag
	fmt.Println("aaaaaaa")

flag:
	fmt.Println("bbbbbbb")
	var aaa interface{} = person{"xzy", 123}
	r, ok := aaa.(person)
	fmt.Println(r, ok)
	testrt, _, _ := testret()
	fmt.Println(testrt)
}

func testret() (int, int, int) {
	return 1, 2, 3
}
