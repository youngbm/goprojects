package main

import "fmt"

func main() {
	v := struct {
		x int
		y int
	}{
		10,
		100,
	}

	type myStruct struct {
		x int
		y int
		z int
		j string
	}
	var x myStruct
	x = myStruct{1, 2, 4, "fuck"}

	fmt.Println(x.x)
	fmt.Println(x.y)
	fmt.Println(x.j)
	fmt.Println(v)
	fmt.Print("fuck")
}
