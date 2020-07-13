package main

import (
	"flag"
	"fmt"
	"os"
)

type ErrorHandling int

const (
	ContinueOnError ErrorHandling = 1 << (iota * 2) // 重置为0，下一行自增
	_                                               // 可以跳过
	ContinueOnError1
	ContinueOnError2
	ContinueOnError3
	ContinueOnError4
)

func usage() {
	fmt.Fprintf(os.Stderr, `
    This is usage page:
	-a --action  input your action
	`)
	flag.PrintDefaults()
}

var (
	h    bool
	v, V bool
	a    string
	l    string
)

func main() {

	//fmt.Println(ContinueOnError)
	//fmt.Println(ContinueOnError1)
	//fmt.Println(ContinueOnError2)
	//fmt.Println(ContinueOnError3)
	//fmt.Println(ContinueOnError4)
	fmt.Println("--main-start--")

	flag.BoolVar(&h, "h", false, "Help")
	flag.StringVar(&a, "a", "", "input your action.")
	flag.StringVar(&l, "l", "low", "input your level.")

	flag.Usage = usage
	flag.Parse()

	if h {
		flag.Usage()
	}

	if a == "" {
		fmt.Println("No action")
	} else {
		fmt.Println("Action:", a)
	}

	fmt.Println("Level:", l)
}
