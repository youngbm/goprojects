package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("--main-start--")
	s := make([]int, 5)
	fmt.Printf("len: %d cap:%d \n", len(s), cap(s)) // 5, 5

	s = append(s, 0, 1, 2, 3, 4)
	fmt.Printf("len: %d cap:%d \n", len(s), cap(s)) // 10, 10
	fmt.Println(s)                                  // [0-10]

	fmt.Println()
	// Next
	s = make([]int, 0)
	fmt.Printf("len: %d cap:%d \n", len(s), cap(s)) // 0, 0

	s = append(s, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("len: %d cap:%d \n", len(s), cap(s)) // 当len增长,先填满容量，再根据自身操作系统提供的算法增加
	fmt.Println(s)

	// Next
	sum, err := funcMui(1, 2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sum)

	// new  make  区别
	fmt.Println()
	myNew := new([]int) // 返回类型的指针
	fmt.Println(*myNew)
	fmt.Printf("%d-%d\n", len(*myNew), cap(*myNew))

	myMake := make([]int, 0) // 返回类型的引用，类型有限制:slice、map 和 channel.,可以预先配置容量
	fmt.Println(myMake)
	fmt.Printf("%d-%d\n", len(myMake), cap(myMake))
}

//func funcMui(x, y int) (sum int, error) { //  sum  又是int类型, 又是error类型
func funcMui(x, y int) (sum int, err error) { // 正确
	sum = x + y //  上面已经声明了变量能直接使用
	err = nil   //  上面已经声明了变量能直接使用
	return x + y, nil
}
