package main

import "fmt"

func f(i int){
    if i--; i == 0 {
        return
    }
    f(i)
}

func main(){
    var val int
    println(&val)
    f(6000) // 由于多次递归会导致协程重新申请栈内存
    println(&val) // 重新申请栈内存导致变量地址更改  println 慎用
    fmt.Println(&val) // fmt.Println 使变量发生逃逸，逃到堆内存，堆内存地址不会改变
}
