package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start...")
	go f1()
	time.Sleep(1 * time.Second)
	fmt.Println("end...")
}

func f1() {
	defer func() {
		recover()
	}()
	//defer recover() //  wrong  syntax
	panic("f1")
	go f2()
}

func f2() {
	defer func() {
		recover()
	}()
	//defer recover()  //  wrong  syntax
	panic("f2")
}
