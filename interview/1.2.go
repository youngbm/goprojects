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
	fmt.Println("I am in f1")
	go f2()
	panic("f1") //  发生panic,就会去defer了,后面不执行
	fmt.Println("I am in f1 end")
}

func f2() {
	defer func() {
		recover()
	}()
	fmt.Println("I am in f2")
	//defer recover()  //  wrong  syntax
	panic("f2")
}
