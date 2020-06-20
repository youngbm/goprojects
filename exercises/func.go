package main

import (
	"fmt"
)

func change(s ...string) {
	s[0] = "Go"
}

func main() {
	welcome := []string{"fuck", "hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
