package main

import "fmt"

func watShadowDefer(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		ret = 10 // 11
		//ret := 10 //  短变量坑死人  100
		defer func() {
			ret = ret + 1 // 这里也是用的 if里面的短变量
		}()
	}
	return
}

func main() {
	fmt.Println("--main-start--")
	r := watShadowDefer(50)
	fmt.Println(r)
}
