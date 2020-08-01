package main

import "fmt"

const (
	azero = iota
	aone  = iota
)

const (
	infor = "msg"
	bzero = iota
	bone  = iota
)

func main() {
	fmt.Println("--main-start--")
	count := 0
	for i := range [256]int{} {
		//n, m := byte(i), int8(i)
		m, n := byte(i), int8(i)
		if n == -n {
			fmt.Println(n)
			fmt.Printf("%d == %d\n", n, -n)
			count++
		}
		if m == -m {
			fmt.Println(m)
			fmt.Printf("%d == %d\n", m, -m)
			count++
		}
	}
	fmt.Println(count) // 内存溢出(-128)和0的时候相等

	// Next
	fmt.Println(azero, aone) // 0, 1
	fmt.Println(bzero, bone) // 1, 2
}
