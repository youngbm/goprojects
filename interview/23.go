package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]    //{2,3}
	s2[1] = 400     //  用索引能够直接修改 最初的切片值
	fmt.Println(s1) // {1,2,400}
	fmt.Println(s2) // {2,400}

	s2[0] = 11 //这里会导致s1修改
	s2[1] = 22

	s2 = append(s2, []int{}...)           //append一个空数组，不扩容, s1 s2还是同底数组
	s2 = append(s2, []int{11, 22, 33}...) //底层数组扩容了生产新的数组给s2

	fmt.Println("s1", s1) //s1 {1,2,400}
	fmt.Println("s2", s2) //s2 {2,400,5,6,7}

	s2[0] = 1111 //如果扩容, s2不会导致s1修改; 如果不扩容s2 和 s1还是同一个底层数组
	s2[1] = 2222
	fmt.Println("s1", s1)

	// Next
	if a := 1; false {
	} else if b := 2; false {
	} else {
		fmt.Println(a, b) // 1 2
	}
}
