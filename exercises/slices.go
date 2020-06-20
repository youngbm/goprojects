package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = [2]int{1, 3}
	fmt.Println(x)

	var y = [...]int{10, 10, 11}
	fmt.Println(y)

	var z = make([]string, 100)
	//for i:=0; k, v := range z {
	for i := 0; i < 100; i++ {
		z[i] = strconv.Itoa(i + 1)
	}
	fmt.Println(z)
}
