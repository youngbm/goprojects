package main

import "fmt"
import "time"
import "runtime"

func main(){
    var ch chan int

    go func(){
        fmt.Println("123") 
    }()

    go func(){  // 这个会马上退哦
         ch = make(chan int, 1)
         ch <- 1
         //time.Sleep(time.Second*1000)
    }()

    go func(ch chan int) {
         time.Sleep(time.Second)
         <-ch
    }(ch)

    for i:= 0 ; i < 10; i++ {
        time.Sleep(time.Second)
        fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine()) // 2  一个是主进程的协程, 一个是line21, 等待读
    }
    //c := time.Tick(1*time.Second)
    //for range c {
    //    fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
    //} 
}
