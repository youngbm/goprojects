package main

import "fmt"

func main() {
	fmt.Println("--main-start--")
	for i := 0; i < 10; i++ {
		//loop:
		fmt.Println(i)
	}
	//goto loop goto 不能跳转到其他函数或者内层代码

	// Next
	x := []int{0, 1, 2}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			print(v) // 3次2
		}()
		y[i] = &v
	}

	print(*y[0], *y[1], *y[2])
}
