package main

import "fmt"

func main() {
	fmt.Println("--main-start--")
	flag := true
	if flag {
		fmt.Println("flag")
	}

	if !!flag {
		fmt.Println("!!flag")
	}

	if flag == true {
		fmt.Println("flag==true")
	}

	// Next
	defer func() {
		fmt.Println(2, recover()) // 2 2
	}()

	defer func() {
		defer func() {
			fmt.Println(1, recover()) //1 1 嵌套defer, 执行也会嵌套分层
		}()
		panic(1)
	}()

	defer recover()
	panic(2)
}
