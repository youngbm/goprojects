package main

import "fmt"

func main() {
	fmt.Println("--main-start--")
	defer func() {
		recover() // 撤回panic
		func() {
			recover() // 这里会优先调用， 则不生效
		}()
	}()
	// recover() 必须是要在panic之后才能生效
	panic("fuck")

	fmt.Println("------END------")

	// Next
}
