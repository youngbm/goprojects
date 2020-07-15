package foo

import (
	"fmt"
)

func init() {
	fmt.Println("bar: --I am function init in pakcage foo--")
}

func SayBar() {
	fmt.Println("sayBar")
}
