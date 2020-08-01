package main

import "fmt"

type T []int

func F(t T) {

}

func main() {
	fmt.Println("--main-start--")
	var q []int //  []int 和 T  能转化, 至少一个Unnamed Typ
	F(q)        // 不同类型也能转化  吹啊
}
