package main

import "fmt"

func sum(n string, s []int, c chan int) {
	sum := 0
	fmt.Println(n, s)
	for _, i := range s {
		fmt.Println(n, i)
		sum += i
	}
	c <- sum
}

func main() {
	fmt.Println("vim-go")
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	c := make(chan int)
	fmt.Println(b[len(b)/2:], b[:len(b)/2])
	go sum("f ->", b[len(b)/2:], c)
	go sum("e ->", b[:len(b)/2], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
