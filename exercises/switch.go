package main

import "fmt"

func main() {
	switch {
	case 1 < 10:
		fmt.Println("1  lt 10")
		fallthrough
	case 1 < 100:
		fmt.Println("1  lt 100")
		fallthrough
	case 1 < 1000:
		fmt.Println("1  lt 1000")
	}
}
