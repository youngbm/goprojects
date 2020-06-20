package main

import (
	"fmt"
	"goprojects/exercises/employee"
)

func main() {
	fmt.Println("main")
	e := employee.New("行政", "孙子家", 27, 10000.0)
	fmt.Println(e)
	salary := e.Zgz(10000.00)
	fmt.Println(e)
	fmt.Println(salary)
}
