package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("--main-start--")
	var myb byte
	fmt.Scanf("%c\n", &myb)

	b := new(bytes.Buffer)
	err := b.WriteByte(myb)
	if err != nil {
		panic(err)
	} else {
		newmyb, _ := b.ReadByte()
		fmt.Printf("%v\n", string(newmyb))
	}
}
