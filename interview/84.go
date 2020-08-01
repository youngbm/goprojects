package main

import "fmt"

func main() {
	fmt.Println("--main-start--")
	a := [3]int{1, 2, 3}
	s := a[1:2] // {2}

	s[0] = 11         // {11}
	s = append(s, 12) // 容量没超过低数组, 还是更改低数组
	s = append(s, 13) // 新的数组切片
	s[0] = 21
	fmt.Println(a) //  {1,11,12}
	fmt.Println(s) //  {21,12,13}
}
