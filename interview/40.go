package main

import "fmt"

//func Stop0(stop <-chan bool) { //cannot close receive-only channel 只能读不能关
//	close(stop)
//}

func Stop1(stop chan<- bool) { //cannot close receive-only channel
	close(stop)
}

func Stop2(stop chan bool) { //cannot close receive-only channel
	close(stop)
}

func main() {
	var ch = make(chan int, 1)
	ch <- 1
	select {
	case ch <- 100:
		fmt.Println("100")
	case v1 := <-ch:
		fmt.Println(v1)
	}

	var i = 2
	switch i {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	default:
		fmt.Println("100")
	}

	// Next
	//定义只读的channel
	//read_only := make(<-chan bool)
	////可同时读写
	//var ch0 = make(chan bool)
	//Stop1(write_only)

	//定义只写的channel
	write_only := make(chan<- bool)
	Stop1(write_only)
	read_write := make(chan bool)
	Stop2(read_write)
	read_write = make(chan bool)
	Stop2(read_write)

	// Next
	//s := new(Show) // 返回指针，并没有创建空间
	s := Show{} // 返回指针，并没有创建空间
	fmt.Println(s.Param)
	//s.Param["day"] = 2 // 没分派空间报错异常, 指针不支持索引
}

type Param map[string]interface{}

type Show struct {
	Param
}
