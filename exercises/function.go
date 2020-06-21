package main

import "fmt"

func sum(i, j int) {
	fmt.Println("func sum", i+j)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println("func a")
	a := func() {
		fmt.Println("hello world  the first class  function!")
	}
	a()
	fmt.Printf("%T\n", a)

	fmt.Println("func b")
	func() {
		fmt.Println("Func b the first class  function!")
	}()

	c := func(f func(int, int)) {
		f(5, 6)
		fmt.Println("高阶函数！", f)
	}
	k := func() {
		fmt.Println("func sum")
	}
	s := sum
	fmt.Println(k)
	c(s)
}
