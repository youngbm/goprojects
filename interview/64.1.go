package main

import "fmt"

func main() {
	fmt.Println("--main-start--")

	defer func() {
		fmt.Print(recover()) //  这里recover会捕获1,并打印
	}()

	defer func() {
		defer fmt.Print(recover())
		panic(1)
	}()

	defer recover() // recover() 必须在 defer() 函数中调用才有效
	panic(2)        // 直接进入defer
	fmt.Println("END")
}
