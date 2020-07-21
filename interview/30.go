package main

import "fmt"

func f(n int) (r int) {
	defer func() {
		r += n
		//recover()
	}()

	var f func()

	defer f() // f() 该函数还没被定义, 后面才定义的, 异常
	f = func() {
		r += 2
	}
	//defer f()

	return n + 1 // return r & return n+1 所以 r = n+1 = 4
}

func main() {
	fmt.Println(f(3))

	// Next
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a { //  相当于开辟了一个函数,副本执行 range
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("r= ", r)
	fmt.Println("a= ", a)
}
