package main

import (
	"fmt"
	"time"
)

type error interface {
	Error() string
}

type myError struct {
	When time.Time
	What string
}

func (e *myError) Error() string {
	return fmt.Sprintf("My defination err: Time: %s , Msg: %s\n", e.When, e.What)
}

//返回接口， 即返回接口的值，即返回方法的接收者 myError
func run(a []int) error {
	var e *myError
	if len(a) == 0 {
		e = &myError{
			time.Now(),
			"I do not work!",
		}
	} else {
		fmt.Println("Get a right parameter!\n")
	}
	return e
}

func main() {
	fmt.Println("vim-go")
	var a []int
	a = append(a, 1)
	err := run(a)
	if err == nil {
		fmt.Println("Get error:", err)
	}
}
