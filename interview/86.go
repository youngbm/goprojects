package main

import "fmt"

const (
	//八进制数以 0 开头，
	//十六进制数以 0x 开头，所以 Decade 表示十进制的 8
	Century = 100
	Decade  = 010 // 8  8进制
	Year    = 001 // 1  8进制
)

func main() {
	fmt.Println("--main-start--")
	n := 43210
	fmt.Println(n/60*60, n%60*60) // 从左往右 依次计算, 数据类型不变; %不能操作 float

	// Next
	fmt.Println(Century + Decade*2 + 2*Year) // 118
}
