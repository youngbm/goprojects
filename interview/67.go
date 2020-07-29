package main

import "fmt"

type T struct {
	n int
}

func main() {
	fmt.Println("--main-start--")
	ts := [2]T{}
	for i := range ts[:] {
		switch t := &ts[i]; i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Println(t.n, " ")
		}
	}
	fmt.Print(ts)
}
