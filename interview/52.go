package main

import "fmt"

type X struct{}

func (x *X) test() {
	fmt.Println(x)
}

func main() {
	var a *X // nil
	a.test() // 打印nil

	var b = X{}
	b.test() // &{}

	//X{}.test() // 结构体不可寻址,不能一步到位

	//  Next
	//v := x["two"]
	//v, ok:= x["two"]
}
