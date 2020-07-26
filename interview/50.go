package main

import "fmt"

// [Important]  数组索引访问都是可以改值
func main() {
	isMatch := func(i int) bool {
		switch i {
		case 1: // 直接跳过, case 完成程序会默认 break
			fmt.Println("fuck")
			fallthrough
		case 2:
			return true
		default:
			return false
		}
		return false
	}

	fmt.Println(isMatch(1)) //false
	fmt.Println(isMatch(2)) //true
}
