package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println("--main-start--")
	r := strings.NewReader("1234567890")
	r.Seek(2, io.SeekStart)
	p := make([]byte, 10)
	r.Read(p)
	fmt.Println(string(p))
}
