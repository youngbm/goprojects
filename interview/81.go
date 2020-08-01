package main

import "fmt"

type TimesMatcher struct {
	base int
}

func NewTimesMatcher(base int) *TimesMatcher {
	return &TimesMatcher{base: base}
}

func main() {
	fmt.Println("--main-start--")
	var a []int = nil
	//a, a[0] = []int{1, 2}, 9 // err  先等号左边的索引表达式和取址表达式，接着计算等号右边的表达式
	fmt.Println(a)

	// Next
	/*
		 空指针 未被初始化的指针
		 野指针 指向未知的内存地址空间
		var p *int
		*p = 0x042058080 // 删除的对象或未申请访问受限内存区域的指针
		fmt.Println("*p", *p)
	*/
	pp := NewTimesMatcher(3)
	fmt.Println(pp)
}
