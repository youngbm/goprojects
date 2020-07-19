package main

import "fmt"

func rt1() (int, int) {
	return 1, 2
}

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

type A interface {
	ShowA() int
}
type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	var x0 int
	x0, _ = rt1()
	fmt.Println(x0)

	var x1 int
	x1, y1 := rt1()
	fmt.Println(x1, y1)

	//var x2 int
	//x2, y2 = rt1() //y2  undefind
	//fmt.Println(x2, y2)
	//x, _ := rt1()  // no new variables on left side of :=   没有新变量不用:=

	// Next
	fmt.Println(increaseA())
	fmt.Println(increaseB())

	var a A = Work{3} // 定义了接口
	//var a = Work{3} // 定义了结构体
	fmt.Println(a.ShowA()) // 2 + 10

	jj := a.(Work) // 把 a的接收者赋给 jj 接口名.(接收者类型)
	fmt.Println(jj.ShowA())
	fmt.Println(jj.ShowB())

	var k = Work{3} // 定义了接口
	fmt.Println(k.ShowA())
	fmt.Println(k.ShowB())
}
