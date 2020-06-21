package main

import (
	"fmt"
)

func main() {
	a := make(map[string]int)
	a["a"] = 123

	delete(a, "a")
	fmt.Println(a["a"])

	v, ok := a["a"]
	fmt.Println(v, ok)

	fmt.Println(a)
}
