package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("--main-start--")
	f, err := os.Open("70.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	// Next
	ff()
	fmt.Println("end")
}

func ff() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover:  %#v\n", r)
		}
	}()
	panic(1) // 这里就返回了,进去 defer
	fmt.Println("ff 1 - 2")
	panic(2)
}
