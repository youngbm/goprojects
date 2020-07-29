package main

import "fmt"

type Ta struct {
	name string
}

func main() {
	fmt.Println("--main-start--")

	//fmt.Println('abc'+'123')  // 字符类型相加, abc为字符串非字符
	fmt.Println("abc" + "123") //  字符串类型相加
	//fmt.Println('123'+"abc")  // 字符加字符串, abc为字符串非字符
	op := fmt.Sprintf("abc%d", 123) // 格式化
	fmt.Println(op)

	// Next
	const (
		x = iota // 置零，后面是递增
		_        // 跳过
		y
		z = "zz" // 赋值
		k        // 同上
		p = iota // 重置为索引值，但是数值不会归0
	)
	fmt.Println(x, y, z, k, p)

	//  Next
	// [Important] nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量

	//var a = nil      //  错误类型 use of untyped nil
	var a *int = nil //  改为指针可以打印
	fmt.Println(a)

	var b interface{} = nil
	fmt.Println(b)

	// 字符不能为nil, 是为空才对， 同理int stuct  等也不能为nil
	var c *string = nil // 改为指针就能打印
	var d *Ta = nil     //
	fmt.Println(c)
	fmt.Println(d)

	var e error = nil
	fmt.Println(e)
}
