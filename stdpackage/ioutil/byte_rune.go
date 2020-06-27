package main

import "fmt"

func main() {
	fmt.Println("--main-start--")

	// rune  int32 用来装 utf-8编码, 可以统一占位的长度
	var astring = "我是大帅哥abc"
	fmt.Println(len(astring))         // 18
	fmt.Println(len([]byte(astring))) //  return 18
	fmt.Println(len([]rune(astring))) // return 8
}
