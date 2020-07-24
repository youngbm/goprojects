package main

import "fmt"

func main() {
	var ch chan int // 这才是nil
	if ch == nil {
		fmt.Println("ch is nil")
	}
	//ch <- 10  //恐慌
	//<-ch		//恐慌

	ch1 := make(chan int) // 这个不为nil
	close(ch1)
	//ch1 <- 10          //恐慌
	<-ch1              // 读出 0
	fmt.Println(<-ch1) //channel 接收数据，如果缓冲区中为空，则返回一个零值

	fmt.Println("end")

	// Next
	const i = 100 // 常量
	var j = 123

	//fmt.Println(i, &i) // 常量在栈, 里面不能取址
	fmt.Println(j, &j)

	// Next
	fmt.Println("Q3\n")
	intmap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	v, err := GetValue(intmap, 3)
	fmt.Println(v, err)

}

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "exist", true
	}
	return "", false
	//return nil, false  //  string不能retrun nil
}
