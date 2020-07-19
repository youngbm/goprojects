package main

import "fmt"

type Person struct {
	age int
}

func main() {
	person := &Person{30}

	defer fmt.Println(1, person.age)

	defer func(p *Person) { // 这里是原来的结构体被压进去了
		fmt.Println(2, p.age)
	}(person)

	defer func() {
		fmt.Println(3, person.age)
	}()

	person = &Person{300} // 这里是和上面2个不同的结构体, 只是指针变量名一样而已
	//person.age = 300 // 如果直接修改原来机构体的值, 2输出300
}
