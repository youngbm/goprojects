package main

import "fmt"

type Foo struct {
	bar string
}

func main() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}

	s2 := make([]*Foo, len(s1))

	for k, v := range s1 {
		tt := v
		//s2[k] = &v //  v 变量地址不变，只是改变值
		s2[k] = &tt //  创建中间变量
	}

	fmt.Println(s1[0], s1[0], s1[2])
	fmt.Println(s2[0], s2[0], s2[2]) //  &{C} &{C} &{C}

	// Next
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}

	counter := 0
	for k, v := range m { // m不是副本, 注意k,v作用域
		if counter == 0 { // 一开始0未必是A
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("Counter is ", counter) //  2 or 3  map不按照顺序来
}
