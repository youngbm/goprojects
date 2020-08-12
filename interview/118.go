
package main

import "fmt"

const (
    _ = iota
    c1 int = (10*iota)
    c2 //  = (10*iota) 逻辑和上一个相对
    d = iota
)

func main(){
    m := make(map[int]int, 10)
    //fmt.Println(cap(m)) // 这个不能cap
    fmt.Println(len(m))
    c := make(chan int, 10)
    fmt.Println(cap(c))

    // Next
    fmt.Println(c1,c2,d)
}
