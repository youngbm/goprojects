package main

import (
	"fmt"
)

type abc struct {
	test int
	//arr  []int  // 不可以比较的类型   map, slice 和 function
	//myfunc func() int
}

func main() {

	sn1 := struct {
		age  int
		name string
		x    abc
	}{age: 11, name: "xzy"}

	sn2 := struct {
		age  int
		name string
		x    abc
	}{age: 11, name: "xzy"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	// Next
	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11,
		m: map[string]string{"a": "1"},
	}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11,
		m: map[string]string{"a": "1"},
	}
	//if sm2 == sm1 { //invalid operation: sm2 == sm1 (struct containing map[string]string cannot be compared))
	if sm2.age == sm1.age {
		fmt.Println("sm1 == sm2")
	}
}
