package main

import "fmt"

func f1() (r int) {
	defer func() { // defer 是 函数返回前 会调用
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t++
	}()
	return t
}

func f3() (r int) {
	defer func(r int) { // 这里执行的是副本
		fmt.Println("in f3", r)
		r = r + 5
		fmt.Println("in f3", r)
	}(r)
	return 1
}

func f4() (r int) {
	defer func(r *int) { // 这里执行的是地址能改值
		fmt.Println("in f4", *r)
		*r = *r + 5
		fmt.Println("in f4", *r)
	}(&r)
	return 100 // 这个就是 r,  r = 100 然后调用 defer
}

func main() {
	fmt.Println("f1", f1()) // 1
	fmt.Println("f2", f2()) // 5

	fmt.Println()
	fmt.Println("f3 main", f3()) // 5

	fmt.Println()
	fmt.Println("f4 main", f4()) // 5
}
