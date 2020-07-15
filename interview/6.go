package main

import "fmt"

type name struct {
	name string
}

type MyInt1 int   // 声明新的的类型
type MyInt2 = int // 只是取个别名, 而且还能用int直接给MyInt2赋值

func main() {
	fmt.Println("--main-start--")

	var p *name
	p = &name{name: "xzy"} // & 对刚刚声明的结构取地址
	//  *p 这个不能赋值, 为空指针
	//*p = name{name: "xzy"} // panic: runtime error: invalid memory address or nil pointer dereference

	fmt.Println(p.name)
	fmt.Println((*p).name) // (*p) 先取值

	var i *int
	a := 1   // 不能直接对数字寻址
	i = &(a) // 如果a的值变了那么i的值也会更改
	fmt.Println(i)
	fmt.Println(*i)

	a = 2 // i 为a的引用,a为2, i变为2
	fmt.Println(i)
	fmt.Println(*i)

	//Next,  data type transform
	var j int = 1000
	//var i1 MyInt1 = j //  类型不同不能直接复制,下面强制类型转换
	var i1 MyInt1 = MyInt1(j)
	var i2 MyInt2 = j

	fmt.Println(i1, i2)

	// Next
	aa := []int{1, 2}
	fmt.Printf("%v\n", aa)
	ap(aa)
	fmt.Printf("%v\n", aa)
	app(aa)
	fmt.Printf("%v\n", aa)
	// 传地址
	myap(&aa)
	fmt.Printf("%v\n", aa)
	// 传指针
	var cc *[]int
	cc = &aa
	myap(cc)
	fmt.Printf("%v\n", aa)
}

func ap(a []int) {
	a = append(a, 100) // a为副本不修改值
}

func app(a []int) {
	a[0] = 100000 // a 为副本, 但是 a[0]却是地址,所以a会修改第一位的值
}

func myap(a *[]int) { // 地址或者指针就会修改值
	*a = append(*a, -100, 0) // a为副本不修改值
}
