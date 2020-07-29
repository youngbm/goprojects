package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s1 := s[:0:2]                     // 1 limit the capacity
	fmt.Println(s1, len(s1), cap(s1)) // 0 1

	// Next
	m := make(map[string]int)
	m["abc"] = 123
	//if v := m["c"];  v != nil {
	if v, ok := m["c"]; ok {
		fmt.Println(v, "Got m[c]")
	} else {
		fmt.Println(v, "Not Got m[c]")
	}

}
