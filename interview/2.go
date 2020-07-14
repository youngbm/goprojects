package main

import "fmt"

// for range 循环的时候会创建每个元素的副本，而不是元素的引用，
// 所以 m[key] = &val 取的都是变量 val 的地址，
// 所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3.
func main() {
	fmt.Println("--main-start--")

	s := []int{0, 1, 2, 3, 4}
	m := make(map[int]*int) //  value为int指针
	var z int
	fmt.Println(&z)
	fmt.Println(z)
	fmt.Println("---------------------")

	for k, z := range s {
		fmt.Println(&z)
		fmt.Println(z)
		//  m[k] 存的是z的变量地址，最后变量地址指向数组的最后一个;
		//  所以打印出来value都是值为切片最后一位的地址
		m[k] = &*&z
		// 正确写法, 申请另外一个新的变量地址,进行赋值
		temp := z
		m[k] = &temp
	}

	fmt.Println()
	for k, v := range m {
		fmt.Println(k, *v)
		fmt.Println(k, v)
	}
}
