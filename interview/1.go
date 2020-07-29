package main

import (
	"fmt"
)

// 1. 当协程遇到panic时，会遍历该协程的defer并执行.
// 2. 如果在defer执行的过程遇到recover则停止panic，返回recover继续向下执行，
// 3. 没有遇到recover，遍历完defer,然后想stderr抛出panic信息。
// 4. 协程中 的panic 必须在该协程内 defer, 不然会直接中断

func defer_call() {
	defer func() {
		fmt.Println("1")
		recover() // 同协程的,一个panic只需要一个recover
	}()
	defer func() {
		fmt.Println("2")
		//recover()
	}()
	defer func() {
		fmt.Println("3")
		//recover()
	}()

	panic("get panic") // 这个最后执行， 如果有recover 就不执行
	//log.Fatalln("get panic") // 直接退出
}

func main() {
	fmt.Println("--main-start--")
	defer_call()
	fmt.Println("--main-end--")
}
