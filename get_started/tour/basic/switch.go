package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Try switch........")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("OS: %s\n", os)
	}
}
