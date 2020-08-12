package main

import "fmt"

func main(){
    a := []int{0, 1}
    fmt.Printf("%v", a[len(a):]) // a[2:] index 越界的切怕切片为空[]
}
