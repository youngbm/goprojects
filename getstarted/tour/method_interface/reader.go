package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	c := make([]byte, 8)
	//var b []byte
	//c := append(b, 0, 0, 0, 0, 0, 0, 0, 0)
	r := strings.NewReader("I am a new NewReader !!!!")
	for {
		n, err := r.Read(c)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, c)
		fmt.Printf("b[:n] = %q\n", c[:n])
		if err == io.EOF {
			break
		}
	}
}
