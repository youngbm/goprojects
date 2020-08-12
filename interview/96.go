package main

import "fmt"

func get() []byte {
	raw := make([]byte, 1000)
	fmt.Println(len(raw), cap(raw), &raw[0]) //  1000 1000 同样的地址
	return raw[:3]                           // 返回了原来的底数数组
}

func main() {
	fmt.Println("--main-start--")
	raw := get()
	fmt.Println(len(raw), cap(raw), &raw[0]) //  3 1000 同样的地址
}
