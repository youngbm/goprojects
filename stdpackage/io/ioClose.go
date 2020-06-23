package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("--main-start--")
	file, err := os.Open("/tmp/123")
	if err != nil {
		fmt.Println("I got this  err")
		panic(err)
	}
	defer file.Close()
}
