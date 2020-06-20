package main

import (
	"fmt"
	"time"
)

type author struct {
	name string
	age  int
}

func (a author) fullName() {
	fmt.Printf("I  am %s, %d years old \n", a.name, a.age)
}

func main() {
	fmt.Println("main")
	time.Now()

}
