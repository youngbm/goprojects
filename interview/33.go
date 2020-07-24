package main

import "fmt"

//C. for 循环支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break，可以选择中断哪一个循环；
//D. for 循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量；
func main() {
	i := 1
	s := []string{"A", "B", "C"}

	i, s[i-1] = 2, "Z" // 多重赋值  类似 python
	fmt.Println(i, s)
}
