package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("--main-start--")
	//bf := bytes.NewBufferString("1234567890")
	const maxInt = int(^uint(0) >> 1)
	//const maxInt = int(^uint(0) >> 1)
	fmt.Println(maxInt)
	fmt.Println(^uint8(0) >> 1)
	fmt.Println(^(^uint8(0) >> 1))

	s := []byte("abcdefgddd")
	seq := []byte("d")
	//Count
	c := bytes.Count(s, seq)
	fmt.Println(c)
	// upper lower
	fmt.Println(bytes.ToUpper(s))
	fmt.Println(bytes.ToLower(s))
	// equal
	d := []byte("abcdefgdd")
	fmt.Println(bytes.Equal(s, d))
	// trim
	fmt.Println("before:", d)
	fmt.Println("after:", bytes.Trim(s, "d"))
	//split
	s = []byte("a1b1c1d1ef1g1d1dd")
	seq = []byte("1")
	for k, v := range bytes.Split(s, seq) {
		fmt.Println(k, string(v))
	}
	//join
	fmt.Printf("%s\n", bytes.Join(bytes.Split(s, seq), []byte("|||")))

	fmt.Println("\n")

	//new
	b := bytes.NewBufferString("\n")
	fmt.Println(b)
	//b.buf = []byte("123123")
	fmt.Println(b.String())
	fmt.Println(b.Len())
	for i := 0; i < 10; i++ {
		r := strings.NewReader("abc123\n")
		b.ReadFrom(r)
	}
	fmt.Println(b, "\n")
}
