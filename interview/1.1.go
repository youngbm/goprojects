package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fmt.Println("start....")
	for w := 0; w < 1; w++ {
		go f1()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("end....")
	fmt.Printf("%T\n", os.Stdout)
	io.Writer()
}

func f1() {
	defer func() {
		recover()
	}()
	fmt.Println("f1")
	panic("pf1")
	go f2()
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
