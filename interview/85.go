package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("--main-start--")
	//rimRight() 会将第二个参数字符串里面所有的字符拿出来处理，只要与其中任何一个字符相等，便会将其删除。
	//想正确地截取字符串，可以参考 TrimSuffix() 函数
	fmt.Println(strings.TrimRight("ABBA", "BA")) // 空
	fmt.Println(strings.TrimSuffix("ABBA", "BA"))

	//  Next
	var src, dst []int
	src = []int{1, 2, 3}
	dst = make([]int, len(src)) // 修复
	copy(dst, src)              // dst 内存还没有申请,所以还是[]
	fmt.Println(dst, src)
}
