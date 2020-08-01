package main

import (
	"fmt"
)

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func main() {
	fmt.Println("--main-start--")
	// 同级文件的包名不允许有多个

	// Next
	d1 := data{"one"}
	d1.print()

	//var in printer =  data{"two"}  //struct 和 interface不能赋值
	in := data{"two"}
	in.print()
}
