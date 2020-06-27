package foo

import "fmt"

//func init() {
//	fmt.Println("This is my first package --- [utilset]]!\n")
//}

func init() {
	fmt.Println("I am init!!!")
	fmt.Println("I am in file  SayHello.go !")
}

func SayHello() {
	fmt.Println("vim-go")
	fmt.Println("This is my first package!")
	fmt.Println("Hello world!!!!")
}
