package main

import "fmt"

func testConvertableParameter(x int, y ...int) {
	fmt.Println(y)
	for k, v := range y {
		fmt.Printf("Idx -> %d, Value -> %v \n", k, v)
	}
}

func main() {
	testConvertableParameter(1, 2, 3, 4, 5, 6, 7, 8)
	testConvertableParameter(2, 3, 4, 5, 6, 7, 8)
	testConvertableParameter(3, 4, 5, 6, 7, 8)
	testConvertableParameter(4, 5, 6, 7, 8)
}
