package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("--main-start--")
	//NewBufferSize
	nr := strings.NewReader("abc123456")
	r := bufio.NewReader(nr)
	fmt.Println(r.Size())
	fmt.Println(r.Peek(3))
	fmt.Println(r.ReadString('s'))
	fmt.Println(r.ReadString('b'))
	fmt.Println(r.Buffered())
}
