package main

import (
	"fmt"
)

func recoverFunc() {
	if ret := recover(); ret != nil {
		fmt.Println("recover from ", ret)
	}
}

func fullName(lastname *string, firstname *string) string {
	defer recoverFunc()
	if lastname == nil {
		panic("lastname is null")
	} else if firstname == nil {
		panic("fristname is null")
	}
	fmt.Println("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	return fmt.Sprintf("%s %s", *lastname, *firstname)
}

type person struct {
	lastname  string
	firstname string
}

func main() {
	fmt.Println("vim-go")
	var xzy = person{"zhiyang", "xie"}
	fmt.Println(fullName(&xzy.firstname, &xzy.lastname))
	fmt.Println("\n\n\n\n")
	var v string
	v = "HHHHHHHHHHHHHHHHHHH"
	fmt.Printf("START|%v|END", fullName(nil, &v))
}
