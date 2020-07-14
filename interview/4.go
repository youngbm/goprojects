package main

import "fmt"

func append_new() {
	fmt.Println("--main-start--")
	list := new([]int)       // 返回是一个 切片指针，并非切片
	*list = append(*list, 1) // 第一个必须是切片，返回也是切片
	fmt.Println(*list)
}

func append_arr() {
	a := []int{1, 2, 3} // 都可以动态扩容
	b := []int{4, 5}
	c := make([]int, 0) // 都可以动态扩容

	c = append(c, 4, 5, 6, 7, 8, 9) // append  第二个参数必须是 单个值，或者多个单值
	//a = append(a, b)                // append  第二个参数必须是 单个值，或者多个单值, b 为切片必须报错
	c = append(c, c...) // append  第二个参数必须是 单个值，或者多个单值, b为切片必须报错, 用b... 打散传入
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

var (
	size     = 1024 // var 关键字出现 不是 := , 而是 =
	max_size = 1024 * 2
	//max_size := 1024 * 2  // 错误写法
	a = 1
	b = 2
)

//c := 3  // 全局声明变量 必须在 var 后面
var d = 4

func print_globalvar() {
	fmt.Println(size, max_size)
	fmt.Println(a, b)
	c := 3 // 函数体里面才能用 :=
	fmt.Println(c)
	fmt.Println(d)
}

func main() {

	append_new()

	append_arr()

	print_globalvar()
}
