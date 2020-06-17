package main

import (
	"fmt"
)

//type Stringer interface {
//	String() string
//}

type person struct {
	name string
	age  int
}

type IPaddress [4]byte

func (p person) String() string {
	return fmt.Sprintf("I am super %s , %d years old!", p.name, p.age)
}

func (ip IPaddress) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	fmt.Println("vim-go")
	var xzy = person{"fjh", 31}
	fmt.Println(xzy)
	var myIP = IPaddress{10, 0, 1}
	fmt.Println(myIP)
}
