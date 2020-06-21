package main

import (
	"fmt"
	"time"
)

//func main() {
//	t := time.Now()
//	switch {
//	case t.Hour() < 12:
//		fmt.Println("Time: AM")
//	case t.Hour() > 12:
//		fmt.Println(t.Hour)
//		fmt.Println("Time: PM")
//	default:
//		fmt.Println("Now 12 o'clock!!!")
//	}
//}

func main() {
	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() > 12:
		fmt.Println("Good afternoon.")
		fmt.Println(t.Hour())
	default:
		fmt.Println("Good evening.")
	}
}
