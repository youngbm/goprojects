package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start....")
	for w := 0; w < 2; w++ {
		go f1()
	}
	time.Sleep(10 * time.Second)
	fmt.Println("end....")
}

func f1() {
	defer func() {
		recover()
	}()
	fmt.Println("f1")
	go f2()
	panic("pf1")
}

func f2() {
	defer func() {
		fmt.Println("df2")
		//fmt.Fprintlni(i, "df2")
		recover()
	}()
	fmt.Println("f2")
	panic("pf2")
}
