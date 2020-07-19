package main

import "fmt"

func main() {
    var s1 []int      // only defind variable s1
    var s2 = []int{}  // = is here
    fmt.Println(s1, s2)
    if s1 == nil {
        fmt.Println("s1 nil")
    } else {
        fmt.Println("s1 not nil")

    }
    if s2 == nil {
        fmt.Println("s2 nil")
    }else {
        fmt.Println("s2 not nil")
    }

    // Next
    i := 65
    fmt.Println(string(i), fmt.Sprintf("%d",i))
}