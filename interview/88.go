package main

import "fmt"

func main() {
	fmt.Println("--main-start--")
	m := make(map[string]int)
	m["foo"]++ //  就是那么酷
	fmt.Println(m["foo"])

}
