package main

import "fmt"

func main() {
	fmt.Println("--main-start--")

	fmt.Println(3^2+4^2 == 5^2)  // true   ^为异或操作符
	fmt.Println(6^2+8^2 == 10^2) // false
}
