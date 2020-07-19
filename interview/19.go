package main

import "fmt"

type Person struct {
	age int
}

func main() {

	//person := &Person{28}
	person := Person{28}

	defer fmt.Println(1, person.age) //28 这里不是调用匿名函数，直接传28了

	defer func(p *Person) {
		fmt.Println(2, p.age) //29  传入指针能修改值
	}(&person)

	defer func(p Person) {
		fmt.Println(3, p.age) //28  创建副本不能修改值
	}(person)

	defer func() {
		fmt.Println(4, person.age) //29 匿名函数形成闭包引用，输出29. 这里没传变量进来 是闭包
	}()

	person.age = 29
}
