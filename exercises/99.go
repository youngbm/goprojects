package main

import "fmt"

func main() {
	fmt.Println("--main-start--")
	for i := 0; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if i*j < 10 && j != 1 {
				fmt.Printf("%d*%d = %2d  ", j, i, i*j)
			} else {
				fmt.Printf("%d*%d = %d  ", j, i, i*j)
			}
		}
		fmt.Println()
	}
}
