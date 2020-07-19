package main

import "fmt"

func myPrint(i *int){     
    fmt.Println(*i)
}

func main(){
    s := "hello"
    //s[0] = 'x'    // varchar is constant type
    fmt.Println(s)
    // Next
}