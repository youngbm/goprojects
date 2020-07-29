package main

import "fmt"

type Orange struct {
	quantity int
}

func (s *Orange) String() string {
	return fmt.Sprintf("%#v", s.quantity)
}

func main() {
	fmt.Println("--main-start--")
	o := Orange{123}
	fmt.Println(o)  // {123} 直接打印{123}
	fmt.Println(&o) // 123  这个才能正常调用方法String()
}
