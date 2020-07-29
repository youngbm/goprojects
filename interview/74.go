package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("--main-start--")

	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		os.Exit(1)
	}()

	//for {   //  这样太耗费cpu了
	//}
}
