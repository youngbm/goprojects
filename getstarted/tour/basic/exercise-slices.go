package main

import (
	"fmt"
	//"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var ret = make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		temp := make([]uint8, dx) // 这里一定要重新声明变量
		for j := 0; j < dx; j++ {
			temp[j] = uint8((i + j) / 2)
		}
		ret[i] = temp
	}
	fmt.Println(ret)
	return ret
}

func main() {
	fmt.Println("Main...")
	Pic(10, 10)
	//pic.Show(Pic(10, 10))
}
