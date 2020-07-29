package main

import "fmt"

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i) //  i 变量地址不变.  i 为最后值 i++ -> 2
		})
	}
	return funs
}

var fn = func(i int) {
	print("x")
}

func main() {
	fmt.Println("--main-start--")
	funs := test()
	for _, f := range funs {
		f()
	}

	// Next
	fn := func(i int) {
		print(i)
		if i > 0 {
			fn(i - 1) // 这里调用全局的函数fn;并不是递归
		}
	}
	fn(10) // 10x
}
