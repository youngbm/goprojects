package main

import "fmt"

type User struct {
	name string
}

func (u *User) SetName(name string) {
	u.name = name
	fmt.Println(u.name)

}

type Employee User

func main() {
	fmt.Println("--main-start--")
	e := new(Employee)
	//e.SetName("Jack") // Employee和User不同类型
	fmt.Println(*e)
}
