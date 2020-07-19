package main

import "fmt"

func myPrint(i *int){  // 如果传入副本则打印出来的的是调用前的副本
    fmt.Println(*i)
}


type People struct {
    name string
}

type Teacher struct {
    People  // 这里不是 p People 继承了People
}

func (p People) ShowA(){
      fmt.Println("p ShowA")
      p.ShowB()
}

func (p People) ShowB(){
      fmt.Println("p ShowB")
}

func (t Teacher) ShowB(){
      fmt.Println("t ShowB")
}


func main(){
    var i = 5
    defer myPrint(&i) // 如果传进去是非指针or地址，则是当前的副本，打印的是5
    i =  10
    defer myPrint(&i) // 如果传进去是非指针or地址，则打印的是10
    i = i + 10
    myPrint(&i)

    // Next
    var t Teacher
    t.ShowA()
}