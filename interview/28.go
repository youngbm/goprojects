package main

import "fmt"

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	fmt.Println(p)
}

func main() {

	fmt.Println([...]int{1} == [1]int{1}) //true 数据类型同,  能比较
	fmt.Println([2]int{1} == [2]int{1})   //true 数据类型同,  能比较
	// [...]int{1}  等于 [1]int{1}

	//fmt.Println([...]int{1} == [2]int{1}) // 数据类型都不同,  不能比较
	//fmt.Println([]int{1} == []int{1})     // []不可以比较的类型, 只能和nil比较, [1]int可以比较

	// Next

	fmt.Println(p) //   p 为 nil
	//fmt.Println(*p) //  *int 指向为空 异常了

	var err error
	p, err = foo() // [Important]  := 和 = 的区别啊  大哥大大大
	//p, err := foo() // [Important] := 作用域是局部的，导致p为新变量,覆盖的全局变量p
	if err != nil {
		fmt.Println(err)
		return
	}

	bar()
	fmt.Println(p)
	fmt.Println(*p)
}
