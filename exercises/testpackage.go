package main

import (
	"fmt"
	. "io"
	mrand "math/rand"
	"time"

	"goprojects/exercises/foo"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(time.Now())
	fmt.Println(mrand.Intn(100))
	fmt.Println(EOF)
	foo.SayHello()
}
