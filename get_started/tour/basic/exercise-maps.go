package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("exercise....")
	ret := make(map[string]int)
	a := "I a m a s m a r t o l d b o y !"
	b := strings.Split(a, " ")
	for _, v := range b {
		_, ok := ret[v]
		if ok {
			fmt.Println(v)
			ret[v] += 1
		} else {
			ret[v] = 1
		}
	}

	fmt.Println(ret)
	fmt.Println(ret["I"])
	v, ok := ret["I"]
	fmt.Println(v, ok)
}
