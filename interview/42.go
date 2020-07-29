package main

import "fmt"

var gvar int //  全局变量可以申明不用

func main() {
	//var one int
	//two := 2
	//var three int
	//three = 3

	func(unused string) { //  形参不用也行
		fmt.Println("Unused!")
	}("What ?")

	// Next
	b := ConfigOne{}
	c := &ConfigOne{}
	fmt.Println(b)
	fmt.Println(c)

	// Next
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)

	for k, _ := range a {
		if k == 0 {
			a = append(a, 6, 7)
		}
		r = append(a, 6, 7)
	}

	fmt.Println(a)
	fmt.Println(r)
}

type ConfigOne struct {
	Daemon string
}

//func (c *ConfigOne) String() string { // 只对 *ConfigOne生效
func (c ConfigOne) String() string { // 对 b, c都生效
	return fmt.Sprintf("right")
}
