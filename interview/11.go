package main

import "fmt"

func main() {
    var a = make([]int, 1)
    //var b = make(map[int]int) // map can be cap()
    var ch = make(chan int, 1)
    fmt.Printf("%d-%d", cap(a), cap(ch))
}