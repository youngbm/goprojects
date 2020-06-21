package main

import (
	"fmt"
	"strings"
)

func main() {
	var cp = [][]string{
		[]string{"1", "2", "3"},
		[]string{"4", "5", "6"},
		[]string{"7", "8", "9"},
	}

	for i := 0; i < len(cp); i++ {
		fmt.Printf("%s\n", strings.Join(cp[i], "|"))
	}
}
