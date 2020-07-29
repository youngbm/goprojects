package main

import "fmt"

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}

func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Println(*s, elem)
	return s // 传指针,最后的值会变
}

func main() {
	fmt.Println("--main-start--")
	s := NewSlice()
	defer s.Add(1).Add(2) // 先执行一波, 只剩下一个调用给defer执行
	s.Add(3)
}
