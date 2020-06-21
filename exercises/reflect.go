package main

import (
	"fmt"
	"reflect"
)

type i interface {
	/* TODO: add methods */
}

type student struct {
	name string
	age  int
}

type teacher struct {
	name    string
	subject string
}

func getName(i interface{}) {
	t := reflect.TypeOf(i)
	k := t.Kind()
	//fmt.Println("type", t)
	//fmt.Println("kind", k)
	if k == reflect.Struct {
		fmt.Println("Number of fields is", t.NumField())
	}
	fmt.Println(i)

	v := reflect.ValueOf(i)
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
	}
}

func main() {
	fmt.Println("vim-go")
	i := 10
	fmt.Printf("Value: %v; Type: %T\n", i, i)
	s := student{"xzy", 10}
	t := teacher{"mn", "math"}
	getName(s)
	fmt.Println("\n\n")
	getName(t)
}
