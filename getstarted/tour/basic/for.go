package main

import "fmt"

func main() {
	var sum, i int // i 不在这里声明,而在for中声明，在后面不能用
	for i = 0; i < 100; i++ {
		sum += i
		fmt.Println(i, sum)
	}
}
