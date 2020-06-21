package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, c chan int) {
	if t.Left != nil {
		Walk(t.Left, c)
	}
	c <- t.Value
	if t.Right != nil {
		Walk(t.Right, c)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int, 10), make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println(Same(t1, t2))
}
