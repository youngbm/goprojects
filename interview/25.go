package main

import "fmt"

type myInt int

//func (i int) printInt() {  // 不能用常规类型来作为接受者
func (i myInt) printInt() {
	fmt.Println(i)
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (s Student) Speak(think string) (talk string) {
	if think == "speck" {
		talk = "speck"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var i myInt = 123
	i.printInt()

	// Next
	var peo People = Student{} // 初始化结构，以及接收者
	think := "speck"
	fmt.Println(peo.Speak(think))
}
