package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	reader := strings.NewReader("I am a just strings!")
	d := make([]byte, 80)
	p, n := reader.ReadAt(d, 0)
	if n != nil && n != io.EOF {
		fmt.Println(n)
		panic(n)
	}
	fmt.Println(string(d), p, n)

	file, err := os.Create("io2-create.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	file.WriteString("1234567890abcdefg\n")
	a, err := file.WriteAt([]byte("1234567890"), 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}
