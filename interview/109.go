package main

import "fmt"
import "sort"

type T struct {
    v int
}

func main(){
    //a := make([]T, 5)
    a := []T{{1}, {3}, {10}, {100}, {2}}

    // sort.Slice 对复杂的类型进行排序, 
    sort.Slice(a, func(i, j int) bool {return a[i].v > a[j].v}) // i.v和j.v
    fmt.Println(a)
}
